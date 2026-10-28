package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 string }
type GenS5 struct { F0 string; F1 uint; F2 int8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32, p2 uint) uint64 {
    var v6 int64 = -1
    for fi7 := 0; fi7 < 2; fi7++ {
        switch -305 {
        case -2147483648:
            fmt.Println(p0, p0)
        }
        v8 := GenS1{F0: 1}
        _ = fi7
        _ = v8
    }
    if false {
        if ((true && false) || true) {
            var v9 float64 = float64((-2.25 / 1.0))
            fmt.Println(v9, p1)
            _ = v9
        } else {
            var v10 string = ""
            v11 := GenS2{F0: -128, F1: (p0 - (p0 & (^(p0))))}
            var v12 int64 = (v6 * v6)
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS3{F0: -1}
    var v14 uint16 = 862
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v6
    return 18446744073709551615
}
func genFn1() uint64 {
    var v15 uint32 = 2147483647
    fmt.Println(v15)
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v15, v15)
        for fi17 := 0; fi17 < 13; fi17++ {
            switch "foo bar" {
            case "\n\t":
                v18 := []int{(-(fi17)), (fi17 << ((1) & 63))}
                _ = v18
            case "":
                v19 := []int{}
                _ = v19
            }
            _ = fi17
        }
        _ = fi16
    }
    var v20 uint64 = 42
    v20 = (v20 + genFn0(uint8(v15), -65, 266))
    _ = v15
    _ = v20
    return v20
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{270, 2147483648}
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 bool = false
    if v5 {
        _ = len(v3)
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    if len(v2) > 0 {
        v2[0] = 744
    }
    v21 := GenS4{F0: 2147483647, F1: 127, F2: "a"}
    v5 = true
    v22 := GenS1{F0: 1}
    fmt.Println(v5)
    fmt.Println(v5)
    v5 = (-2.25 < 0.0)
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v5)
        if v5 {
            v5 = v5
            v5 = (((42 < 1000) || (true || (((3.14159 + 3.14159) * (-2.25 * 3.14159)) == ((-2.25 / 1e10) * 0.0)))) || (fi23 >= (-(fi23))))
        }
        _ = fi23
    }
    if true {
        var v24 bool = (("\n\t" + ("" + "foo bar")) < "")
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 string = ""
            var v27 int64 = 42
            if v24 {
                delete(v3, fi25)
                _ = len(v2)
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        if (18446744073709551615 <= 42) {
            fmt.Println(v5, v5)
        }
        _ = v24
    }
    var v28 int = 0
    v2 = append(v2, 32)
    v28 += (v28 * (^(v28)))
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            v5 = true
            _ = fi30
        }
        var v31 uint64 = 196
        var v32 int = fi29
        _ = fi29
        _ = v31
        _ = v32
    }
    v2 = append(v2, 2147483648)
    var v33 float32 = -2.25
    if v5 {
        var v34 float64 = 1e10
        if (!false) {
            v35 := GenS2{F0: 127, F1: 0}
            _ = v35
        } else {
            if (!(1 == 0)) {
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 uint8 = 255
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v39 := v1[625]
                        _ = fi38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 3; fi40++ {
                        var v41 string = "\n\t"
                        var v42 uint16 = 623
                        v43 := GenS5{F0: (v41 + ((v41 + (v41 + "\n\t")) + v41)), F1: 938, F2: 116}
                        _ = fi40
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi36
                    _ = v37
                }
                var v44 bool = v5
                _ = v44
            }
            v45 := GenS6{F0: 1}
            _ = v45
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v34
    } else {
        if ((v5 || v5) && (2 < 18446744073709551615)) {
            var v46 uint8 = 53
            fmt.Println(v46, v33, v33)
            _ = v46
        } else {
            if ((!(!v5)) && ((-32768 != 42) && (!(false || v5)))) {
                v47 := GenS4{F0: 549, F1: 165, F2: "\n\t"}
                if (float32(1e10) > v33) {
                    var v48 int = (-(v28))
                    _ = v48
                } else {
                    var v49 int64 = -9223372036854775808
                    var v50 int8 = -128
                    var v51 bool = v5
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = v47
            } else {
                var v52 uint64 = 2
                v53 := v4["foo bar"]
                _ = v52
                _ = v53
            }
            v3[1] = 2147483648
        }
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
