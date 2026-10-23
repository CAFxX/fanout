package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint8) uint64 {
    p0 -= p0
    if true {
        p0 = p0
        if ((255 > p0) || (4294967295 > 2)) {
            v7 := []int{9223372036854775807, 4294967295}
            var v8 uint16 = 0
            var v9 bool = true
            _ = v7
            _ = v8
            _ = v9
        }
    } else {
        p0 += (p0 >> ((0) & 63))
        v10 := GenS1{F0: (p0 + (^(p0))), F1: -1}
        _ = v10
    }
    var v11 bool = false
    var v12 string = ("\n\t" + "")
    _ = p0
    _ = v11
    _ = v12
    return 18446744073709551615
}

func main() {
    var v1 string = ("foo bar" + ("a" + "a"))
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := []string{v1, "hello", "foo bar"}
    var v5 uint = 18446744073709551615
    v6 := []uint32{2147483648}
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v1)
        _ = fi13
    }
    v1 = (v1 + v1)
    fmt.Println(v5)
    v5 -= v5
    fmt.Println(v5, v1)
    v1 = (v1 + v1)
    v5 = (v5 / (((v5 << ((v5) & 63))) | 1))
    if true {
        var v14 float32 = (0.1 - 1e10)
        _ = v14
    } else {
        if true {
            var v15 uint64 = 433
            if false {
                v16 := GenS2{F0: uint16(v5)}
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 uint16 = 42
                    fmt.Println(fi17, v5, v15)
                    v5 = 1
                    _ = fi17
                    _ = v18
                }
                var v19 string = ("foo bar" + "")
                _ = v16
                _ = v19
            } else {
                for fi20 := 0; fi20 < 11; fi20++ {
                    delete(v3, 725)
                    _ = fi20
                }
            }
            _ = v15
        } else {
            if (-870 != 42) {
                switch 632 {
                case 127:
                    var v21 bool = true
                    _ = v21
                }
                var v22 int16 = 1
                var v23 int32 = -494
                _ = v22
                _ = v23
            }
            if (18446744073709551615 != 2) {
                var v24 bool = false
                v1 = ""
                fmt.Println(v5, v24, v5)
                _ = v24
            } else {
                var v25 int64 = -1
                _ = v25
            }
        }
        switch (v1 + "\n\t") {
        case "foo bar":
            if (56 >= 174) {
                if len(v4) > 0 {
                    v4[0] = ""
                }
                v5 = v5
                v5 += (v5 | v5)
            } else {
                fmt.Println(v1, v5)
            }
        case "":
            v26 := v3[334]
            _ = v26
        }
        var v27 uint64 = 4294967295
        _ = v27
    }
    if ((v1 + "") == ("hello" + v1)) {
        v5 += v5
        for fi28 := 0; fi28 < 2; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                v1 = (v1 + "hello")
                if (!((!(v1 != ("" + (("\n\t" + v1) + ("hello" + v1))))) && true)) {
                    v3[1] = 65535
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    if true {
                        v30 := GenS3{F0: 932, F1: "\n\t"}
                        v31 := GenS4{F0: fi28, F1: 0}
                        var v32 int16 = 1000
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                } else {
                    v6 = append(v6, 165)
                }
                fmt.Println(v5, v5)
                _ = fi29
            }
            var v33 int64 = 9223372036854775807
            _ = fi28
            _ = v33
        }
    }
    v5 = 659
    v34 := GenS5{F0: 2147483647}
    var v35 int16 = -1
    var v36 int16 = v35
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
