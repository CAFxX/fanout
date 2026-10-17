package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint16, p2 string) uint {
    v5 := []int{}
    p1 = p1
    p1 += (p1 << ((801) & 63))
    var v6 uint16 = (p1 % (((^(p1))) | 1))
    if (!(544 <= 65535)) {
        p1 += v6
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return uint(p0)
}
func genFn1() uint8 {
    v7 := []int{0}
    var v8 int8 = 42
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 bool = (!(!true))
    var v2 int = 2147483647
    var v3 uint = 255
    var v4 int8 = 0
    if v1 {
        var v9 string = "foo bar"
        if ((-(v4)) > (^(v4))) {
            if true {
                var v10 bool = v1
                _ = v10
            } else {
                v11 := GenS1{F0: -731, F1: 2147483647}
                fmt.Println(v2)
                v4 += 100
                _ = v11
            }
            var v12 string = v9
            v2 = 32768
            _ = v12
        } else {
            v13 := GenS1{F0: int32(v3), F1: (v2 % ((v2) | 1))}
            v14 := GenS2{F0: 32767}
            switch ("hello" + v9) {
            case "\n\t":
                if v1 {
                    var v15 int = v2
                    _ = v15
                }
            case "hello":
                v16 := GenS3{F0: 13}
                _ = v16
            }
            _ = v13
            _ = v14
        }
        if (((!v1) || false) || ((-420 <= -803) || (0 <= 42))) {
            v17 := GenS2{F0: 1}
            var v18 float32 = -2.25
            v19 := GenS2{F0: 2}
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v9
    }
    if v1 {
        for fi20 := 0; fi20 < 14; fi20++ {
            switch (fi20 * 32768) {
            case -9223372036854775808:
                for fi21 := 0; fi21 < 2; fi21++ {
                    if ((v3 << (((v3 - v3)) & 63)) >= v3) {
                        switch fi20 {
                        case 2147483647:
                            if false {
                                v22 := []int{(-(fi21)), v2, v2, (^(v2))}
                                var v23 int64 = 1
                                if ((!false) || (929 >= 65535)) {
                                    var v24 string = "hello"
                                    v2 = 4294967295
                                    _ = v24
                                } else {
                                    var v25 int16 = 808
                                    v3 -= v3
                                    var v26 uint32 = 256
                                    _ = v25
                                    _ = v26
                                }
                                v27 := GenS1{F0: 42, F1: fi21}
                                var v28 uint = v3
                                _ = v22
                                _ = v23
                                _ = v27
                                _ = v28
                            } else {
                                v29 := []int{fi20, fi21}
                                var v30 uint32 = 626
                                v31 := GenS1{F0: 0, F1: v2}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        case 128:
                            var v32 bool = v1
                            _ = v32
                        case 13:
                            v33 := []int{fi21, v2, (fi21 / (((v2 * fi20)) | 1))}
                            _ = v33
                        default:
                            v34 := GenS1{F0: -2147483648, F1: fi20}
                            _ = v34
                        }
                        v2 = v2
                    } else {
                        v35 := GenS3{F0: 101}
                        _ = v35
                    }
                    _ = fi21
                }
            case 1:
                if (0 == 0) {
                    v36 := []int{32768}
                    _ = v36
                }
            case 963:
                for fi37 := 0; fi37 < 18; fi37++ {
                    v38 := []int{fi37}
                    if v1 {
                        v4 = (v4 | v4)
                        v39 := GenS4{F0: -9223372036854775808}
                        v4 += -1
                        if len(v38) > 0 {
                            v38[0] = (fi37 | (v2 << (((^(v3))) & 63)))
                        }
                        v40 := GenS1{F0: 42, F1: fi20}
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                    _ = v38
                }
            default:
                var v41 int64 = 9223372036854775807
                _ = v41
            }
            _ = fi20
        }
        v2 -= v2
    } else {
        var v42 float32 = (-2.25 / float32(float64(3.14159)))
        if (!(v2 <= 128)) {
            var v43 bool = v1
            v44 := []int{-621, (v2 * v2)}
            var v45 uint32 = 2147483648
            _ = v43
            _ = v44
            _ = v45
        } else {
            var v46 int = (v2 | v2)
            v47 := GenS3{F0: genFn1()}
            _ = v46
            _ = v47
        }
        _ = v42
    }
    v1 = (v1 && (v1 || (-1 >= 42)))
    v2 += v2
    fmt.Println(v2, v4, v3)
    v48 := []int{}
    fmt.Println(v1, v4, v1)
    var v49 int32 = 2147483647
    if len(v48) >= 2 {
        v48 = v48[1:len(v48)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
