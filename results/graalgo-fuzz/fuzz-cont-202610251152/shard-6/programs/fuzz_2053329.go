package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int, p1 uint8) int16 {
    var v4 uint32 = 492
    v5 := GenS1{F0: false, F1: 2, F2: 57}
    var v6 int16 = 1
    var v7 int16 = v6
    var v8 float32 = (float32((-2.25 - 1e10)) / 3.14159)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 995
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 0
    var v3 int8 = 63
    v2 += 1
    v3 += v3
    var v9 float32 = 0.0
    v10 := GenS2{F0: 419, F1: 2, F2: 1}
    if (true || (false && (4294967295 != 0))) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS3{F0: 2}
            v13 := GenS4{F0: 2147483647, F1: 1}
            _ = fi11
            _ = v12
            _ = v13
        }
        switch (v2 - v2) {
        case -9223372036854775808:
            if ((v2 >> ((65535) & 63)) <= (v2 >> ((255) & 63))) {
                v14 := GenS3{F0: 8}
                for fi15 := 0; fi15 < 14; fi15++ {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := []int{fi16, fi16, fi15, fi15}
                        var v18 string = ((("a" + (("a" + "\n\t") + "")) + "foo bar") + "")
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    v19 := []int{}
                    var v20 int16 = 1000
                    _ = fi15
                    _ = v19
                    _ = v20
                }
                _ = v14
            } else {
                v2 = -9223372036854775808
            }
        default:
            var v21 uint8 = 127
            _ = v21
        }
        v22 := GenS3{F0: 42}
        _ = v22
    } else {
        if false {
            if true {
                switch "\n\t" {
                case "\n\t":
                    for fi23 := 0; fi23 < 8; fi23++ {
                        var v24 int16 = -959
                        var v25 int8 = -128
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case "hello":
                    var v26 float32 = v9
                    _ = v26
                }
            } else {
                var v27 uint16 = 0
                v2 = v2
                _ = v27
            }
            v28 := GenS5{F0: v2, F1: 173}
            _ = v28
        }
    }
    if ("\n\t" >= "hello") {
        var v29 uint = 42
        var v30 string = "foo bar"
        _ = v29
        _ = v30
    } else {
        if ((v2 - (v2 / ((v2) | 1))) == v2) {
            for fi31 := 0; fi31 < 20; fi31++ {
                v3 = 100
                switch ((("foo bar" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar") {
                case "":
                    for fi32 := 0; fi32 < 6; fi32++ {
                        switch v2 {
                        case 32768:
                            v33 := GenS6{F0: 2, F1: 42}
                            _ = v33
                        case 922:
                            fmt.Println(fi32, fi31, fi32)
                        case 32767:
                            if false {
                                var v34 float32 = 1e10
                                var v35 uint64 = 0
                                _ = v34
                                _ = v35
                            }
                        }
                        var v36 int16 = 820
                        _ = fi32
                        _ = v36
                    }
                default:
                    switch "\n\t" {
                    case "foo bar":
                        v3 = -128
                    case "\n\t":
                        v37 := []int{(v2 | -572)}
                        _ = v37
                    default:
                        v38 := []int{(v2 + int(v3)), fi31, -431, 1}
                        _ = v38
                    }
                }
                v39 := []int{0, (v2 / (((v2 - -1)) | 1))}
                _ = fi31
                _ = v39
            }
        } else {
            var v40 float32 = (float32(0.0) / -2.25)
            if true {
                _ = len(v1)
                var v41 int = (v2 * v2)
                for fi42 := 0; fi42 < 5; fi42++ {
                    for fi43 := 0; fi43 < 10; fi43++ {
                        if false {
                            var v44 int8 = v3
                            v45 := GenS3{F0: 1}
                            v2 = v41
                            _ = v44
                            _ = v45
                        }
                        var v46 uint32 = 65535
                        var v47 int64 = 42
                        _ = fi43
                        _ = v46
                        _ = v47
                    }
                    var v48 string = ((("hello" + (("foo bar" + "hello") + ("" + "hello"))) + "hello") + "foo bar")
                    _ = fi42
                    _ = v48
                }
                _ = v41
            }
            _ = v40
        }
        var v49 float32 = v9
        _ = v49
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        v3 = (-(v3))
        v51 := []int{127}
        _ = fi50
        _ = v51
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
