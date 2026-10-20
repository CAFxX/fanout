package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint32 = 720
    v2 := make(map[string]int)
    var v3 uint32 = (v1 / ((990) | 1))
    var v4 uint16 = 32767
    if (9 <= -1) {
        fmt.Println(v3)
        var v5 uint16 = 0
        _ = v5
    } else {
        delete(v2, "")
    }
    fmt.Println(v1, v1)
    v6 := []int{}
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            v8 := GenS1{F0: "hello"}
            _ = v8
        }
        var v9 uint16 = 65535
        _ = fi7
        _ = v9
    }
    fmt.Println(v3, v3, v1)
    var v10 string = "a"
    switch (v10 + "a") {
    case "foo bar":
        var v11 float64 = (0.1 - float64(float32((3.14159 * (((3.14159 - 0.0) / float64((1e10 - (1.5 - 1e10)))) / ((-2.25 - 1.5) - 3.14159))))))
        _ = v11
    case "\n\t":
        if false {
            if len(v6) > 0 {
                v6[0] = -2147483648
            }
            if false {
                var v12 uint16 = v4
                _ = v12
            } else {
                _ = len(v2)
                for fi13 := 0; fi13 < 17; fi13++ {
                    var v14 uint64 = 462
                    _ = fi13
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 uint16 = v4
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int32 = 2147483647
                    v19 := GenS2{F0: 32767, F1: 124, F2: v3}
                    if true {
                        v20 := GenS2{F0: 42, F1: v3, F2: 1}
                        fmt.Println(fi17)
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if true {
                    if false {
                        v21 := GenS3{F0: v10}
                        _ = v21
                    } else {
                        _ = len(v2)
                    }
                }
                _ = fi15
                _ = v16
            }
        }
    case "":
        v10 = "foo bar"
    }
    switch -257 {
    case 128:
        switch 32768 {
        case -767:
            var v22 uint = 995
            _ = v22
        case -2147483648:
            fmt.Println(v4, v1)
        case 0:
            var v23 uint = 65535
            _ = v23
        }
    case 0:
        var v24 uint64 = 18446744073709551615
        _ = v24
    default:
        var v25 float64 = 0.1
        _ = v25
    }
    var v26 uint32 = (v3 / (((v1 | (^(v3)))) | 1))
    if false {
        v4 -= v4
        var v27 uint32 = (v26 + v26)
        _ = v27
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v28 := GenS4{F0: 65535}
        for fi29 := 0; fi29 < 6; fi29++ {
            v30 := v2[(((("\n\t" + "") + "a") + "hello") + v10)]
            for fi31 := 0; fi31 < 13; fi31++ {
                if (!false) {
                    fmt.Println(fi31, v3, v3)
                } else {
                    v30 = fi31
                }
                fmt.Println(v10, v10, fi29)
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
        _ = v28
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
