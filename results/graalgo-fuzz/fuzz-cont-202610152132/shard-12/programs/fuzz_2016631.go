package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 bool; F2 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 uint, p2 uint32) int8 {
    v7 := GenS1{F0: uint16(p0), F1: uint16(p0), F2: (^(p2))}
    var v8 int8 = 127
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 uint = p1
        v11 := GenS1{F0: 32767, F1: 126, F2: (p2 << ((p1) & 63))}
        var v12 int64 = 275
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 uint64 = (p0 * uint64(p1))
    var v14 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 int32) uint32 {
    if false {
        var v15 bool = true
        switch 32768 {
        case 2147483647:
            v16 := []int{32767, 32767}
            _ = v16
        default:
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 uint64 = 42
                v19 := []int{}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        if (2147483648 != 864) {
            var v20 uint16 = 0
            v21 := []int{4294967295, 127}
            _ = v20
            _ = v21
        }
        _ = v15
    } else {
        var v22 int32 = p0
        for fi23 := 0; fi23 < 6; fi23++ {
            if (true || (460 <= 259)) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    p0 += 2147483647
                    _ = fi24
                }
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS2{F0: (v22 >> ((442) & 63))}
                    var v27 uint32 = 178
                    if (-9223372036854775808 == -9223372036854775808) {
                        v27 -= v27
                    } else {
                        v28 := []int{4294967295, -2147483648, fi23, (fi25 << ((987) & 63))}
                        if len(v28) > 0 {
                            v28[0] = fi25
                        }
                        v22 += (v22 + 362)
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    v30 := []int{}
                    _ = len(v30)
                    _ = fi29
                    _ = v30
                }
                if (true && (true && (!false))) {
                    if true {
                        v31 := GenS2{F0: 2147483647}
                        var v32 uint = 18446744073709551615
                        var v33 int32 = 1
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        if (true && (!false)) {
                            p0 -= (p0 ^ (v22 - (v22 << ((2) & 63))))
                        }
                        v34 := []int{(^(fi23))}
                        if (true && (false || false)) {
                            _ = len(v34)
                            p0 -= 1
                            if len(v34) > 0 {
                                v34[0] = fi23
                            }
                        } else {
                            v35 := GenS2{F0: v22}
                            v34 = append(v34, 1)
                            var v36 int16 = 32767
                            _ = v35
                            _ = v36
                        }
                        _ = v34
                    }
                }
            }
            var v37 int8 = 42
            var v38 float32 = -2.25
            _ = fi23
            _ = v37
            _ = v38
        }
        v39 := GenS2{F0: v22}
        _ = v22
        _ = v39
    }
    fmt.Println(p0)
    v40 := GenS2{F0: 42}
    v41 := []int{-9223372036854775808}
    var v42 int16 = -145
    _ = p0
    _ = v40
    _ = v41
    _ = v42
    return 1
}

func main() {
    v1 := []int{4294967295, -477, 32767, 315}
    v2 := []uint32{1}
    v3 := make(map[uint32]uint32)
    var v4 bool = (-180 <= 183)
    v5 := make(map[uint32]string)
    v6 := []string{}
    if false {
        v4 = v4
    }
    var v43 int = -2147483648
    v4 = v4
    switch "a" {
    case "":
        var v44 string = "\n\t"
        _ = v44
    case "\n\t":
        fmt.Println(v4, v4)
    }
    fmt.Println(v4)
    v45 := GenS2{F0: 2147483647}
    v43 = v43
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 string = "hello"
        if v4 {
            if len(v6) > 0 {
                v6[0] = (v47 + (v47 + (("\n\t" + v47) + v47)))
            }
            switch v47 {
            case "\n\t":
                v48 := v3[65535]
                _ = v48
            default:
                switch fi46 {
                case -27:
                    for fi49 := 0; fi49 < 17; fi49++ {
                        fmt.Println(fi46)
                        _ = len(v3)
                        v50 := GenS3{F0: 1, F1: false, F2: 100}
                        _ = fi49
                        _ = v50
                    }
                }
            }
        }
        var v51 int16 = -32768
        _ = fi46
        _ = v47
        _ = v51
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    delete(v3, 4294967295)
    fmt.Println(v43, v43)
    v4 = v4
    v52 := GenS4{F0: (v43 * (^(v43)))}
    fmt.Println(v43, v43)
    for fi53 := 0; fi53 < 15; fi53++ {
        delete(v5, 677)
        _ = fi53
    }
    if true {
        switch (("hello" + ("a" + "hello")) + "") {
        case "a":
            v43 = (^(v43))
        case "\n\t":
            for fi54 := 0; fi54 < 8; fi54++ {
                var v55 int32 = 461
                _ = fi54
                _ = v55
            }
        }
        if len(v2) > 0 {
            v2[0] = 2
        }
        var v56 uint8 = 241
        _ = v56
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v52
    _ = v6
}
