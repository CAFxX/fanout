package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint8 }
func genFn0() int16 {
    if (-59 == -2147483648) {
        var v6 uint64 = 239
        v6 -= (^(v6))
        _ = v6
    }
    var v7 float64 = 1e10
    _ = v7
    return -1
}
func genFn1(p0 int32) uint {
    var v8 int32 = p0
    var v9 int16 = 42
    v10 := []int{-9223372036854775808}
    fmt.Println(v8, p0)
    for fi11 := 0; fi11 < 5; fi11++ {
        if (!(!false)) {
            v12 := GenS1{F0: 2147483647, F1: 127, F2: 65535}
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int8 = -97
                if false {
                    if false {
                        fmt.Println(fi11, fi11)
                        if (("\n\t" + "\n\t") > "foo bar") {
                            if len(v10) > 0 {
                                v10[0] = (fi13 - fi11)
                            }
                            v15 := GenS2{F0: v9, F1: 255, F2: -32768}
                            var v16 int = 127
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v15
                            _ = v16
                        } else {
                            if len(v10) > 0 {
                                v10[0] = fi11
                            }
                            var v17 uint64 = 0
                            var v18 bool = ((fi13 <= 127) && (("a" + ("a" + "")) >= "hello"))
                            _ = v17
                            _ = v18
                        }
                        p0 -= -2147483648
                    } else {
                        v19 := GenS2{F0: -1, F1: 42, F2: -32768}
                        _ = v19
                    }
                    var v20 int64 = 42
                    _ = v20
                } else {
                    v21 := GenS1{F0: 42, F1: (v14 ^ (v14 << ((255) & 63))), F2: 2}
                    var v22 int32 = (v8 | (p0 | v8))
                    _ = v21
                    _ = v22
                }
                var v23 string = ((("foo bar" + ("foo bar" + "a")) + "") + "\n\t")
                _ = fi13
                _ = v14
                _ = v23
            }
            _ = v12
        } else {
            if (128 > 129) {
                var v24 int64 = 0
                _ = v24
            } else {
                p0 += p0
                if false {
                    var v25 bool = (42 < 127)
                    v10 = append(v10, fi11)
                    _ = v25
                } else {
                    var v26 int32 = (v8 ^ p0)
                    v8 -= 411
                    _ = v26
                }
                v27 := GenS1{F0: 791, F1: 106, F2: 32767}
                _ = v27
            }
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 755
}

func main() {
    v1 := []uint32{42, 127}
    v2 := []string{"hello", "\n\t", "hello", "a"}
    v3 := []string{("" + "foo bar")}
    v4 := []int{128, 32768, 4294967295, -9223372036854775808}
    v5 := []string{("foo bar" + "\n\t"), "a"}
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint16 = 0
        var v30 int = 42
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: 293, F1: 255, F2: 1000}
    for fi32 := 0; fi32 < 1; fi32++ {
        _ = len(v1)
        var v33 bool = false
        _ = fi32
        _ = v33
    }
    var v34 float32 = 0.1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v34 -= v34
    for fi35 := 0; fi35 < 19; fi35++ {
        v34 = (v34 - (3.14159 - ((float32(0.1) * 3.14159) / v34)))
        if true {
            for fi36 := 0; fi36 < 18; fi36++ {
                v37 := GenS3{F0: 91, F1: 561}
                v34 -= v34
                _ = fi36
                _ = v37
            }
            v38 := GenS4{F0: fi35, F1: (fi35 << ((genFn1(216)) & 63))}
            _ = v38
        } else {
            v39 := GenS5{F0: 127}
            for fi40 := 0; fi40 < 5; fi40++ {
                v34 += v34
                fmt.Println(fi35, fi40)
                _ = fi40
            }
            v34 = -2.25
            _ = v39
        }
        _ = fi35
    }
    switch -2147483648 {
    case 32768:
        v34 = v34
    case 9223372036854775807:
        var v41 int32 = 0
        _ = v41
    case 2147483647:
        for fi42 := 0; fi42 < 4; fi42++ {
            var v43 string = "hello"
            _ = fi42
            _ = v43
        }
    }
    fmt.Println(v34)
    if (!true) {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v34)
        case "":
            var v44 uint = 65535
            _ = v44
        default:
            v34 += float32((float64(3.14159) * -2.25))
        }
        var v45 uint8 = 165
        switch -903 {
        case 0:
            v45 = (^(v45))
        case 958:
            var v46 int = -2147483648
            _ = v46
        default:
            v45 = (v45 << ((genFn1(-2147483648)) & 63))
        }
        _ = v45
    } else {
        var v47 int8 = 119
        _ = v47
    }
    var v48 int8 = 100
    var v49 float32 = 1e10
    var v50 uint16 = 183
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
