package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint8) uint32 {
    v7 := GenS1{F0: 9, F1: "foo bar", F2: (p0 & (p0 ^ (^(p0))))}
    p0 -= p0
    v8 := []int{128, 358}
    switch "" {
    case "foo bar":
        p0 += p0
    case "\n\t":
        v9 := GenS1{F0: 0, F1: "hello", F2: p0}
        _ = v9
    default:
        var v10 string = "a"
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}
func genFn1(p0 uint64, p1 string, p2 int64) uint64 {
    p0 -= p0
    var v11 int8 = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{-501}
        _ = fi12
        _ = v13
    }
    v11 += 100
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (^(p0))
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 int8 = -43
    v5 := []string{"hello"}
    var v6 string = ""
    v14 := GenS2{F0: -367, F1: false, F2: 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((42 >= -25) && true) {
        _ = len(v5)
        var v15 bool = (-1 == 834)
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int16 = 32767
            v17 = 32767
            delete(v3, 65535)
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        fmt.Println(v1, v6)
        v1 -= v1
        var v18 uint64 = (v1 + (v1 * v1))
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 uint = 65535
                v4 -= (v4 / ((v4) | 1))
                var v22 uint = v21
                switch "a" {
                case "foo bar":
                    var v23 uint16 = 163
                    _ = v23
                case "\n\t":
                    var v24 bool = (!true)
                    _ = v24
                case "hello":
                    delete(v3, 2)
                }
                var v25 uint = v22
                _ = fi20
                _ = v21
                _ = v22
                _ = v25
            }
            _ = fi19
        }
    }
    if (0 < (^(v4))) {
        delete(v2, v6)
        var v26 uint8 = 0
        _ = v26
    } else {
        var v27 int16 = -1
        v2["a"] = ("\n\t" + v6)
        v4 = (v4 + -1)
        _ = v27
    }
    delete(v2, (v6 + v6))
    if ((-(v4)) <= 100) {
        for fi28 := 0; fi28 < 14; fi28++ {
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v6)
                var v29 int32 = -1
                _ = v29
            } else {
                v30 := v2[v6]
                _ = len(v3)
                _ = v30
            }
            v1 += (v1 % (((^(v1))) | 1))
            _ = fi28
        }
        v6 = ((v6 + v6) + (v6 + ((v6 + v6) + "\n\t")))
        _ = len(v3)
    }
    var v31 string = "a"
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            switch 498 {
            case 128:
                v5 = append(v5, v6)
            case 0:
                if true {
                    if false {
                        fmt.Println(fi32, v6)
                    } else {
                        delete(v2, "foo bar")
                    }
                }
            default:
                v6 = v6
            }
            v4 -= (v4 ^ (v4 % ((int8(fi32)) | 1)))
            v2[v31] = v31
            v1 -= (^(v1))
            if len(v5) > 0 {
                v5[0] = v31
            }
            v34 := GenS1{F0: (^(v4)), F1: (v31 + (v31 + (v31 + "a"))), F2: 255}
            v35 := GenS3{F0: (^(v1))}
            var v36 int32 = 2147483647
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        for fi37 := 0; fi37 < 17; fi37++ {
            v5 = append(v5, (v31 + (v31 + "a")))
            _ = fi37
        }
        _ = fi32
    }
    switch (v31 + "\n\t") {
    case "\n\t":
        v2[("\n\t" + "a")] = (("\n\t" + (("" + "hello") + v6)) + v31)
    }
    v6 = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
