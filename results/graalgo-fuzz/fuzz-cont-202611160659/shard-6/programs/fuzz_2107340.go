package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 string) int64 {
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 string = p0
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = (((p0 + ((p0 + (("foo bar" + p0) + "a")) + p0)) + p0) + "hello")
        var v7 int = -2147483648
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 2147483648
    _ = p0
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1e10
    var v3 int16 = -774
    if (!true) {
        v9 := v1[528]
        _ = v9
    } else {
        v10 := GenS1{F0: 42}
        _ = v10
    }
    var v11 uint8 = 255
    v11 = (v11 - v11)
    switch "" {
    case "hello":
        v3 = (v3 ^ v3)
    }
    if true {
        if false {
            v1[42] = 0
        } else {
            v12 := GenS1{F0: 2}
            fmt.Println(v11)
            v13 := GenS2{F0: v11, F1: "hello"}
            _ = v12
            _ = v13
        }
    }
    v3 = -1
    v14 := []int{-117}
    if (-9223372036854775808 != 42) {
        if true {
            _ = len(v1)
        } else {
            if true {
                v2 = 3.14159
                if false {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        if (1 <= -1) {
                            v3 += 0
                        } else {
                            v2 += 0.1
                        }
                        for fi16 := 0; fi16 < 3; fi16++ {
                            var v17 int64 = 42
                            _ = fi16
                            _ = v17
                        }
                        _ = fi15
                    }
                    var v18 uint16 = 1
                    if false {
                        v18 += (v18 & v18)
                        fmt.Println(v2, v3, v3)
                    }
                    _ = v18
                }
            } else {
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int32 = 0
                    v21 := GenS1{F0: 0}
                    v22 := GenS3{F0: "", F1: -895}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
        }
        v23 := GenS2{F0: (v11 & 42), F1: "foo bar"}
        v24 := GenS2{F0: (v11 << ((1) & 63)), F1: (("hello" + "a") + "")}
        _ = v23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
}
