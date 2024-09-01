(ns beetleman.leetcode-75.merge-strings-alternately-test
  (:require [beetleman.leetcode-75.merge-strings-alternately :as sut]
            [clojure.test :as t :refer [deftest is]]))

(deftest merge-alternately-test
  (is (= ""
         (sut/merge-alternately "" ""))))
