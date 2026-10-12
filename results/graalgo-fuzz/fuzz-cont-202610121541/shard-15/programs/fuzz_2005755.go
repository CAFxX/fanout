package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int32) int8 {
    v6 := GenS1{F0: 652, F1: 0}
    v7 := GenS1{F0: 18446744073709551615, F1: -1}
    v8 := []int{90, 683, 82, 692}
    var v9 uint64 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -109
}
func genFn1(p0 int8, p1 int16, p2 uint64) int8 {
    var v10 uint32 = 780
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int = fi11
        switch (v12 | v12) {
        case 42:
            v13 := []int{254, (v12 - int(v10)), (fi11 >> ((0) & 63)), int(p2)}
            _ = v13
        case 2147483647:
            var v14 float32 = 1e10
            _ = v14
        default:
            var v15 int32 = -812
            _ = v15
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint64 = (p2 % ((18446744073709551615) | 1))
            v17 += (^(v17))
            p0 = (-(p0))
            _ = fi16
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    v18 := GenS1{F0: 2, F1: 926}
    v19 := GenS1{F0: 2, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v18
    _ = v19
    return p0
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint64 = 0
    var v5 bool = ((true || false) && false)
    v1 = v1
    for fi20 := 0; fi20 < 13; fi20++ {
        v5 = (uint32(v4) != 4294967295)
        var v21 float64 = 1e10
        switch 128 {
        case 32768:
            v4 += 18446744073709551615
        case 439:
            var v22 uint64 = (v4 >> ((1000) & 63))
            _ = v22
        default:
            switch ("\n\t" + "") {
            case "hello":
                for fi23 := 0; fi23 < 16; fi23++ {
                    if v5 {
                        if v5 {
                            delete(v3, (fi23 / ((fi23) | 1)))
                            v24 := GenS1{F0: 255, F1: 42}
                            v4 += 1
                            v5 = v5
                            _ = v24
                        } else {
                            v25 := []int{fi23, 32767, (fi23 / ((fi20) | 1))}
                            _ = v25
                        }
                    }
                    _ = fi23
                }
            case "":
                var v26 bool = true
                _ = v26
            case "foo bar":
                if v5 {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v28 := GenS1{F0: 2, F1: 800}
                        v29 := GenS1{F0: 623, F1: 1}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    v30 := []int{-288, fi20, (^(fi20))}
                    if true {
                        if true {
                            v21 -= float64(float32(float64(((v1 * v1) + v1))))
                        }
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        fmt.Println(v21, v1, v1)
                    } else {
                        v31 := GenS2{F0: 339, F1: 20, F2: 1}
                        var v32 int8 = 42
                        v32 = v32
                        _ = v31
                        _ = v32
                    }
                    _ = v30
                } else {
                    fmt.Println(fi20)
                    if (32767 > 944) {
                        if true {
                            v4 = (v4 >> ((989) & 63))
                            for fi33 := 0; fi33 < 5; fi33++ {
                                delete(v2, "a")
                                fmt.Println(fi33, v4)
                                switch "\n\t" {
                                case "hello":
                                    var v34 uint16 = 65535
                                    _ = v34
                                default:
                                    v35 := GenS1{F0: 411, F1: 138}
                                    _ = v35
                                }
                                var v36 float32 = (v1 - 0.1)
                                fmt.Println(fi20, fi33, v1)
                                v37 := GenS2{F0: 42, F1: 255, F2: 0}
                                _ = fi33
                                _ = v36
                                _ = v37
                            }
                        } else {
                            _ = len(v3)
                            v1 = v1
                        }
                        fmt.Println(fi20)
                        v38 := GenS2{F0: 912, F1: 42, F2: 0}
                        _ = v38
                    }
                    var v39 int16 = -32768
                    _ = v39
                }
            }
        }
        _ = fi20
        _ = v21
    }
    delete(v2, "foo bar")
    v4 = (v4 & (v4 % ((794) | 1)))
    var v40 uint = 65535
    v40 += (v40 + (v40 * 0))
    v3[-9223372036854775808] = "foo bar"
    v3[32767] = "foo bar"
    var v41 uint8 = 2
    fmt.Println(v41, v4, v4)
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS3{F0: 1000, F1: 9223372036854775807, F2: v4}
        for fi44 := 0; fi44 < 4; fi44++ {
            switch -9223372036854775808 {
            case 986:
                v4 -= v4
            }
            v1 = 1e10
            if v5 {
                switch ("" + ("a" + "foo bar")) {
                case "foo bar":
                    v45 := v2["hello"]
                    _ = v45
                case "a":
                    switch "foo bar" {
                    case "foo bar":
                        var v46 string = ((("\n\t" + "\n\t") + ("foo bar" + ("hello" + "\n\t"))) + "hello")
                        _ = v46
                    case "\n\t":
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            var v47 uint8 = 42
                            _ = v47
                        case "hello":
                            delete(v2, "")
                        case "foo bar":
                            v5 = v5
                        default:
                            v5 = (!v5)
                        }
                    default:
                        v48 := GenS4{F0: v4}
                        _ = v48
                    }
                case "":
                    fmt.Println(v41, v4, fi44)
                default:
                    v1 -= 3.14159
                }
            } else {
                var v49 int = (-(fi42))
                if false {
                    fmt.Println(v1)
                    var v50 string = "hello"
                    fmt.Println(v4)
                    _ = v50
                }
                switch (v49 ^ (fi42 & (v49 % ((-773) | 1)))) {
                case 32768:
                    var v51 string = "a"
                    _ = v51
                }
                _ = v49
            }
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    if (!(0 == 515)) {
        v52 := GenS1{F0: (v40 - (v40 >> ((v40) & 63))), F1: 32767}
        var v53 uint16 = 255
        _ = v52
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
