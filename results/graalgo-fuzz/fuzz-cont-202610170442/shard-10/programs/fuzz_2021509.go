package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 string; F2 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 int8 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 int8 = -1
    v3[""] = 65535
    v6 := []int{2147483647, 0}
    fmt.Println(v1, v5)
    fmt.Println(v1)
    if ((v1 & 0) == (v1 / ((1) | 1))) {
        fmt.Println(v2)
    }
    if false {
        v7 := GenS1{F0: 2147483647}
        _ = v7
    } else {
        v8 := GenS2{F0: ("\n\t" + ("hello" + ("\n\t" + "foo bar"))), F1: "a", F2: "a"}
        _ = v8
    }
    v9 := GenS1{F0: 4294967295}
    v1 += v1
    var v10 uint32 = 0
    delete(v3, (("foo bar" + "") + ("" + "foo bar")))
    v11 := GenS1{F0: 2147483647}
    for fi12 := 0; fi12 < 19; fi12++ {
        if false {
            if (1 != 0) {
                var v13 uint16 = v1
                if (false || false) {
                    if (false && false) {
                        var v14 bool = true
                        _ = len(v3)
                        v15 := GenS3{F0: 128, F1: 103, F2: (v5 * (v5 * (v5 & v5)))}
                        _ = v14
                        _ = v15
                    }
                    var v16 string = "a"
                    if (v10 == (v10 * (v10 + v10))) {
                        v13 += 65535
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v1 -= (^(v1))
                            var v18 uint32 = v10
                            _ = fi17
                            _ = v18
                        }
                    }
                    _ = v16
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int = (-(fi19))
                        _ = fi19
                        _ = v20
                    }
                    if len(v6) > 0 {
                        v6[0] = -2147483648
                    }
                }
                delete(v4, ("hello" + "hello"))
                _ = v13
            }
        }
        _ = fi12
    }
    var v21 bool = true
    v5 = 0
    var v22 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
