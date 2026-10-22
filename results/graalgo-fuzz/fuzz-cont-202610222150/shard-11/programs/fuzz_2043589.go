package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint8, p2 uint) int64 {
    v4 := []int{}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 uint64) int32 {
    if (("a" + "a") != "") {
        switch ("hello" + (("a" + "foo bar") + "")) {
        case "hello":
            switch 42 {
            case 32767:
                var v5 uint32 = 2
                _ = v5
            case -217:
                v6 := GenS1{F0: 1, F1: 42}
                _ = v6
            case 42:
                p0 -= (p0 >> ((65535) & 63))
            default:
                var v7 bool = ((1 > 2) && false)
                _ = v7
            }
        case "":
            switch -592 {
            case 128:
                p0 -= (p0 % ((p0) | 1))
            default:
                p0 += (^(p0))
            }
        }
        v8 := GenS1{F0: 1, F1: 127}
        _ = v8
    }
    if true {
        var v9 uint16 = 65535
        switch ("hello" + "") {
        case "a":
            v10 := []int{127}
            _ = v10
        case "\n\t":
            fmt.Println(v9, p0)
        case "":
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int = 127
                _ = fi11
                _ = v12
            }
        default:
            if (0.0 == -2.25) {
                var v13 bool = (9223372036854775807 != 42)
                _ = v13
            }
        }
        _ = v9
    } else {
        if true {
            v14 := []int{-690, 334, 42}
            switch "" {
            case "hello":
                var v15 uint32 = 0
                _ = v15
            case "\n\t":
                var v16 int16 = 42
                _ = v16
            }
            fmt.Println(p0)
            _ = v14
        } else {
            p0 = (p0 >> ((42) & 63))
        }
    }
    switch -2147483648 {
    case 32767:
        var v17 float32 = -2.25
        _ = v17
    case -297:
        fmt.Println(p0)
    case 129:
        switch 1 {
        case 4294967295:
            v18 := GenS2{F0: 39, F1: -1, F2: 154}
            _ = v18
        case 32768:
            v19 := []int{-2147483648, 32767, 2147483647}
            _ = v19
        case -1:
            fmt.Println(p0)
        }
    default:
        if (255 == 725) {
            p0 = (^(p0))
        } else {
            p0 -= 0
            var v20 float32 = 1.5
            _ = v20
        }
    }
    p0 -= 1
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS2{F0: 32767, F1: 1, F2: 0}
        _ = fi21
        _ = v22
    }
    _ = p0
    return -623
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 float32 = 3.14159
    var v3 bool = false
    v23 := GenS2{F0: 65535, F1: 13, F2: 255}
    for fi24 := 0; fi24 < 6; fi24++ {
        v2 += v2
        v2 = (v2 - (0.0 * -2.25))
        _ = fi24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v2, v3, v1)
        v26 := []int{fi25}
        _ = fi25
        _ = v26
    }
    v1 = ""
    v27 := []int{}
    var v28 float32 = (v2 + v2)
    var v29 float64 = 0.1
    v3 = (!v3)
    if true {
        if ((float64((v28 - (-2.25 + 1.5))) - v29) == (v29 - 0.0)) {
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 int32 = 0
                switch (fi30 * (fi30 + fi30)) {
                case 764:
                    var v32 int8 = -1
                    _ = v32
                case 107:
                    var v33 string = v1
                    _ = v33
                case 128:
                    if (-1 >= -1) {
                        if v3 {
                            v27 = append(v27, 965)
                            v34 := GenS3{F0: -9223372036854775808}
                            v2 += float32(0.0)
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 11; fi35++ {
                                fmt.Println(v29, fi35)
                                _ = fi35
                            }
                            var v36 uint16 = 42
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            _ = v36
                        }
                    }
                default:
                    v28 -= ((((v28 - float32((((v29 - v29) * (1.5 * -2.25)) * -2.25))) + 0.1) * 0.0) - v2)
                }
                if len(v27) > 0 {
                    v27[0] = fi30
                }
                _ = fi30
                _ = v31
            }
        }
        v37 := GenS2{F0: 1, F1: 1, F2: 0}
        _ = v37
    }
    fmt.Println(v3, v3, v2)
    var v38 float32 = v28
    switch ((v1 + (v1 + (v1 + "a"))) + (v1 + "\n\t")) {
    case "a":
        switch v1 {
        case "foo bar":
            if len(v27) > 0 {
                v27[0] = 128
            }
        case "hello":
            v2 = v28
        default:
            v38 -= float32(v29)
        }
    case "hello":
        _ = len(v27)
    case "":
        var v39 bool = (!v3)
        _ = v39
    default:
        v28 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
}
