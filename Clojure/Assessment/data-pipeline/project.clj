(defproject data-pipeline "0.1.0-SNAPSHOT"
  :description "Data Pipeline with Web and Database - Clojure Assessment"
  :url "https://github.com/Sanjeev2498/NAPL"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [[org.clojure/clojure "1.12.2"]
                 [ring/ring-core "1.12.1"]
                 [ring/ring-jetty-adapter "1.12.1"]
                 [ring/ring-json "0.5.1"]
                 [compojure "1.7.1"]
                 [org.clojure/java.jdbc "0.7.12"]
                 [org.xerial/sqlite-jdbc "3.45.1.0"]
                 [cheshire "5.12.0"]]
  :main ^:skip-aot data-pipeline.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
