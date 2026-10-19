package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint8; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 60
    var v3 int8 = 127
    var v4 int8 = v3
    v5 := []string{"", "a", "foo bar"}
    v6 := GenS1{F0: 2147483647, F1: (("" + ((("hello" + "") + "") + "\n\t")) + ("\n\t" + ((("a" + "\n\t") + "foo bar") + "hello")))}
    for fi7 := 0; fi7 < 9; fi7++ {
        v3 = (v3 % (((v3 & v3)) | 1))
        _ = fi7
    }
    v3 -= v3
    if (true && (false && (4294967295 < 4294967295))) {
        v8 := GenS2{F0: "foo bar"}
        fmt.Println(v4, v3, v4)
        if (-32768 >= -462) {
            if ((-2.25 * float32(0.1)) > float32(0.1)) {
                if (1 < 2) {
                    fmt.Println(v3)
                    switch (("\n\t" + "a") + "hello") {
                    case "\n\t":
                        v2 += (v2 * v2)
                    case "a":
                        v9 := GenS3{F0: true}
                        _ = v9
                    case "foo bar":
                        v10 := v1[0]
                        _ = v10
                    }
                } else {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        fmt.Println(v2)
                        _ = len(v5)
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 14; fi12++ {
                        v2 = v2
                        v13 := GenS2{F0: "hello"}
                        _ = fi12
                        _ = v13
                    }
                    v4 = v4
                }
            } else {
                if (!true) {
                    if len(v5) > 0 {
                        v5[0] = (("" + (("foo bar" + "") + ("a" + "hello"))) + "a")
                    }
                    delete(v1, 128)
                    var v14 uint64 = 0
                    _ = v14
                }
                v15 := GenS4{F0: 0}
                v4 = v3
                _ = v15
            }
            v3 = v3
        }
        _ = v8
    }
    v16 := GenS5{F0: 1, F1: 2, F2: "\n\t"}
    delete(v1, 2147483647)
    if true {
        var v17 uint8 = 125
        _ = v17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
