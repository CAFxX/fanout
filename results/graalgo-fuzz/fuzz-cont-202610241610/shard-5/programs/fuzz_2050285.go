package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 438
    var v2 float64 = (1.5 / 1.5)
    var v3 int16 = 1
    var v4 int = 0
    v5 := GenS1{F0: "foo bar", F1: (^(v1)), F2: (!true)}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: "a", F1: (^(v1)), F2: (!true)}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 4294967295}
    if true {
        v2 = 0.1
        switch ("a" + "hello") {
        case "a":
            var v9 int32 = 45
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            if (!(!(v3 != v3))) {
                fmt.Println(fi10)
            } else {
                v11 := []int{(fi10 ^ (v4 >> ((18446744073709551615) & 63))), (v4 & v4), -1}
                if true {
                    _ = len(v11)
                    var v12 float64 = v2
                    var v13 int64 = 343
                    _ = v12
                    _ = v13
                } else {
                    _ = len(v11)
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v1 = (^(v1))
                        _ = fi14
                    }
                    _ = len(v11)
                }
                fmt.Println(v4, v4)
                _ = v11
            }
            var v15 int64 = -946
            _ = fi10
            _ = v15
        }
    }
    var v16 int64 = 1
    v17 := []int{(v4 + (v4 % (((v4 - -9223372036854775808)) | 1))), v4, v4}
    var v18 uint = v1
    v19 := GenS1{F0: ("" + "hello"), F1: (^(v1)), F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
