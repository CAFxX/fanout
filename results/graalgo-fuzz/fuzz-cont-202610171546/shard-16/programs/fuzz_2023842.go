package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int64; F1 uint64; F2 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 uint8, p2 int8) uint16 {
    v4 := GenS1{F0: 9223372036854775807, F1: 128}
    v5 := GenS1{F0: p0, F1: 0}
    var v6 int16 = 0
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = (p0 | -1)
        v8 := []int{fi7, fi7, 127, fi7}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 492
}
func genFn1() int32 {
    var v9 uint = 2
    if false {
        v9 = (^(v9))
    } else {
        var v10 uint64 = 0
        if (447 != uint32(v9)) {
            switch -1 {
            case -565:
                fmt.Println(v10)
            case 2147483647:
                v11 := []int{-139, 9223372036854775807, 4294967295}
                _ = v11
            case -9223372036854775808:
                v9 = v9
            default:
                fmt.Println(v9)
            }
            var v12 int = 1
            v13 := []int{4294967295}
            _ = v12
            _ = v13
        } else {
            v14 := []int{}
            _ = v14
        }
        _ = v10
    }
    fmt.Println(v9)
    v15 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
    _ = v15
    _ = v9
    return 1
}

func main() {
    var v1 int8 = -1
    var v2 string = ("a" + "a")
    var v3 string = v2
    var v16 float64 = 1.5
    v17 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
    v2 = v2
    if ((!false) || true) {
        v18 := []int{0}
        if (!true) {
            if (0 >= 32767) {
                v19 := GenS1{F0: 418, F1: 2147483647}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                _ = v19
            }
            var v20 uint8 = 255
            _ = v20
        } else {
            v21 := GenS2{F0: 9223372036854775807, F1: 42, F2: v3}
            v2 = v3
            _ = v21
        }
        var v22 uint64 = 18446744073709551615
        _ = v18
        _ = v22
    } else {
        v2 = ""
        if true {
            for fi23 := 0; fi23 < 5; fi23++ {
                fmt.Println(fi23, fi23)
                var v24 int = fi23
                if (v3 < v3) {
                    v3 = ""
                    fmt.Println(v1)
                }
                _ = fi23
                _ = v24
            }
            var v25 int16 = 1000
            v26 := GenS2{F0: 1, F1: 1, F2: v3}
            _ = v25
            _ = v26
        } else {
            var v27 uint64 = 0
            fmt.Println(v16)
            _ = v27
        }
        v2 = (((v3 + "") + v2) + (v2 + ("a" + "hello")))
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 string = (v3 + v3)
        for fi30 := 0; fi30 < 18; fi30++ {
            if (537 >= 933) {
                if false {
                    v31 := []int{(fi28 >> ((65535) & 63)), (fi30 | fi30), 32768, (fi30 - fi28)}
                    v2 = ("" + "hello")
                    var v32 uint = 18446744073709551615
                    _ = v31
                    _ = v32
                } else {
                    if (true || (-473 > 647)) {
                        v2 = ""
                    } else {
                        switch ((v29 + v3) + v2) {
                        case "foo bar":
                            switch 0 {
                            case 333:
                                v33 := []int{(fi28 ^ (fi30 ^ 4294967295))}
                                _ = v33
                            case 9223372036854775807:
                                var v34 int64 = -9223372036854775808
                                _ = v34
                            case -343:
                                if (false || true) {
                                    if true {
                                        v35 := []int{127, (fi30 & (fi28 >> ((0) & 63))), fi30, 325}
                                        v3 = (("foo bar" + v2) + v29)
                                        _ = v35
                                    }
                                    var v36 bool = false
                                    v36 = v36
                                    var v37 uint = 2
                                    _ = v36
                                    _ = v37
                                } else {
                                    v38 := []int{}
                                    if (true || true) {
                                        v1 += 42
                                    }
                                    fmt.Println(v3, v3)
                                    _ = v38
                                }
                                v39 := []int{4294967295}
                                _ = v39
                            }
                        case "a":
                            if ((^(v1)) > (v1 / ((0) | 1))) {
                                v40 := GenS3{F0: -1}
                                v3 = v2
                                var v41 uint = 4294967295
                                _ = v40
                                _ = v41
                            }
                        case "":
                            fmt.Println(v29, v29, v3)
                        default:
                            v16 = float64(3.14159)
                        }
                    }
                    v3 = v3
                    if (!true) {
                        v42 := []int{fi30, fi30, fi28, fi28}
                        _ = v42
                    }
                }
            }
            fmt.Println(fi28, fi30)
            _ = fi30
        }
        v2 = v2
        _ = fi28
        _ = v29
    }
    v43 := GenS2{F0: 0, F1: 537, F2: v3}
    var v44 bool = true
    v45 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}
