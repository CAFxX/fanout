package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint16) int32 {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 float32 = (0.0 - (float32(1.5) * (((0.0 * ((0.0 * 3.14159) - ((0.0 / 1.0) * 0.1))) + float32((3.14159 + 1.5))) + (((0.1 + 1.5) / 1.0) / 1e10))))
            _ = fi5
            _ = v6
        }
        var v7 uint8 = 2
        _ = fi4
        _ = v7
    }
    p0 += p1
    v8 := []int{-697}
    switch -995 {
    case 2147483647:
        fmt.Println(p1, p0)
    case -2147483648:
        var v9 float64 = float64(float32(1e10))
        _ = v9
    case -376:
        var v10 int64 = -9223372036854775808
        _ = v10
    }
    p0 -= 42
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 4294967295
    v3 := []uint32{uint32(v1)}
    var v11 uint64 = 4294967295
    if false {
        v11 -= (v2 % (((v11 & 148)) | 1))
    }
    v12 := GenS1{F0: -1}
    _ = len(v3)
    var v13 uint16 = 856
    for fi14 := 0; fi14 < 7; fi14++ {
        _ = len(v3)
        _ = fi14
    }
    v15 := GenS2{F0: 2}
    for fi16 := 0; fi16 < 8; fi16++ {
        v13 += (v13 | 170)
        _ = fi16
    }
    fmt.Println(v2)
    v13 = (v13 << ((2) & 63))
    if true {
        for fi17 := 0; fi17 < 1; fi17++ {
            if true {
                if false {
                    fmt.Println(v13)
                    var v18 uint64 = 80
                    if (!true) {
                        v19 := GenS3{F0: ("\n\t" + ("" + "a"))}
                        _ = v19
                    }
                    _ = v18
                } else {
                    v1 += v1
                    v1 -= 42
                    var v20 uint = 255
                    _ = v20
                }
                v21 := GenS2{F0: 65535}
                var v22 uint8 = uint8(v2)
                _ = v21
                _ = v22
            } else {
                if ((float32(float64(1e10)) * 3.14159) != 3.14159) {
                    var v23 float32 = float32(1e10)
                    if ((((((^(fi17)) < fi17) || false) || true) && false) && false) {
                        switch "a" {
                        case "foo bar":
                            v23 = 1.5
                        case "":
                            for fi24 := 0; fi24 < 7; fi24++ {
                                var v25 int16 = 0
                                v26 := GenS4{F0: (v25 + v25)}
                                v1 += 160
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        default:
                            var v27 string = ""
                            _ = v27
                        }
                        if (false && false) {
                            var v28 int8 = 94
                            var v29 int = fi17
                            var v30 int8 = (v28 ^ v28)
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 5; fi31++ {
                            v2 = (v2 | v11)
                            fmt.Println(v2, v13)
                            _ = fi31
                        }
                    }
                    v2 -= 1
                    _ = v23
                } else {
                    v1 += v1
                    if true {
                        var v32 uint32 = 2147483647
                        _ = v32
                    }
                    var v33 uint8 = 42
                    _ = v33
                }
                switch fi17 {
                case 333:
                    if false {
                        fmt.Println(v2, v13, fi17)
                        v13 = (v13 & uint16(v11))
                    }
                }
                v13 -= (^(v13))
            }
            _ = fi17
        }
        var v34 uint8 = 42
        v35 := GenS5{F0: 0}
        _ = v34
        _ = v35
    }
    fmt.Println(v13, v11, v1)
    var v36 int64 = 0
    var v37 int64 = 42
    switch 0 {
    case -9223372036854775808:
        if true {
            var v38 int16 = -787
            v13 = v13
            for fi39 := 0; fi39 < 9; fi39++ {
                var v40 bool = (!true)
                _ = fi39
                _ = v40
            }
            _ = v38
        }
    case -858:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch ("foo bar" + (("a" + ("\n\t" + ((("foo bar" + "\n\t") + "foo bar") + ("foo bar" + ("a" + "hello"))))) + (("foo bar" + "foo bar") + "hello"))) {
    case "":
        switch "a" {
        case "\n\t":
            fmt.Println(v1)
        case "a":
            if len(v3) > 0 {
                v3[0] = 1
            }
        default:
            _ = len(v3)
        }
    case "\n\t":
        var v41 float32 = 3.14159
        _ = v41
    case "foo bar":
        v11 += 42
    default:
        var v42 uint32 = uint32(v2)
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
}
