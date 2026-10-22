package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint16; F1 int64; F2 string }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int8, p1 int, p2 uint16) uint32 {
    p2 = (^(p2))
    p0 = p0
    v6 := []int{p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 996
}
func genFn1(p0 int32) string {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 & fi7), fi7}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: 1000}
        p0 = 247
        _ = fi9
        _ = v10
    }
    _ = p0
    return "\n\t"
}

func main() {
    var v1 uint64 = 934
    var v2 uint = 2
    v3 := make(map[string]string)
    var v4 bool = true
    var v5 uint16 = 1
    v2 -= (v2 << (((v2 >> ((v2) & 63))) & 63))
    for fi11 := 0; fi11 < 5; fi11++ {
        if v4 {
            v2 = v2
            if ((!v4) || (("foo bar" + ("foo bar" + "\n\t")) == "")) {
                if true {
                    v12 := GenS1{F0: 775}
                    switch "foo bar" {
                    case "hello":
                        for fi13 := 0; fi13 < 9; fi13++ {
                            v14 := GenS2{F0: 140}
                            v5 += uint16(v1)
                            v1 = (v1 + (v1 * 4294967295))
                            v15 := GenS1{F0: 1000}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    case "\n\t":
                        v16 := GenS3{F0: -189, F1: ""}
                        _ = v16
                    case "":
                        var v17 float32 = (float32((float64(0.0) - 0.1)) - 0.0)
                        _ = v17
                    }
                    v5 = (v5 | (v5 ^ (^(v5))))
                    _ = v12
                } else {
                    fmt.Println(fi11, v5, v5)
                }
            } else {
                v2 += uint(v5)
                fmt.Println(v4, v1)
            }
            v18 := v3["\n\t"]
            _ = v18
        }
        _ = fi11
    }
    v2 -= 255
    if v4 {
        v1 = v1
        v19 := GenS2{F0: 255}
        v3["a"] = "a"
        _ = v19
    }
    v20 := v3["a"]
    if ((((!true) && (1e10 >= 0.1)) && v4) || (!true)) {
        fmt.Println(v4, v20)
        v5 = v5
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int64 = -842
            var v23 float64 = 0.1
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        switch -807 {
        case 32768:
            v4 = true
        case 127:
            _ = len(v3)
        case 9223372036854775807:
            fmt.Println(v20, v4, v1)
        default:
            var v24 float32 = 1e10
            _ = v24
        }
        if v4 {
            v25 := GenS1{F0: 0}
            _ = v25
        } else {
            v26 := []int{-226}
            v5 -= uint16(v2)
            v27 := GenS4{F0: 255, F1: 9223372036854775807, F2: v20}
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            var v29 uint8 = 1
            _ = fi28
            _ = v29
        }
    }
    if ((v5 <= (^(v5))) || (v1 > v1)) {
        v4 = ((v1 >> (((v2 + 0)) & 63)) < v1)
        if false {
            v20 = ((("hello" + v20) + ((v20 + v20) + "foo bar")) + "")
            switch v20 {
            case "hello":
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 string = "a"
                    v32 := GenS5{F0: 527}
                    v33 := []int{fi30, fi30}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
        } else {
            var v34 int32 = -1
            _ = v34
        }
    }
    _ = len(v3)
    v35 := GenS2{F0: 134}
    v36 := GenS4{F0: v5, F1: 0, F2: "foo bar"}
    for fi37 := 0; fi37 < 10; fi37++ {
        v38 := GenS1{F0: -32768}
        var v39 int = (fi37 >> (((v2 / ((1) | 1))) & 63))
        v40 := []int{(fi37 ^ (v39 / (((v39 << (((^(v2))) & 63))) | 1))), (fi37 / ((0) | 1)), (-(v39))}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
