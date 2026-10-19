package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0(p0 int, p1 bool) uint8 {
    p0 = 1
    if (206 < 224) {
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 10; fi7++ {
                v8 := GenS1{F0: 127}
                fmt.Println(p1)
                p0 = fi7
                _ = fi7
                _ = v8
            }
            p1 = ((("a" >= ("hello" + "hello")) && (!true)) && false)
            if (false && ((0 == -690) || (2147483647 >= 65535))) {
                p1 = (!(797 <= 2))
            }
            _ = fi6
        }
        v9 := []int{456, p0, p0}
        _ = v9
    }
    var v10 int = p0
    p1 = ((3.14159 * (((-2.25 - 0.0) / 0.1) * 0.0)) <= float32(1e10))
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 float64 = float64((float32(1.5) * 3.14159))
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return 28
}

func main() {
    var v1 int32 = 503
    var v2 float64 = float64(3.14159)
    var v3 int16 = -32768
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 65535
    var v13 float64 = (3.14159 * (v2 / float64(1e10)))
    v14 := []int{863, 9223372036854775807}
    var v15 uint32 = uint32(v5)
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            switch (fi16 >> ((926) & 63)) {
            case 42:
                v3 += int16(v5)
            case 32768:
                var v17 int64 = 1
                _ = v17
            case 32767:
                fmt.Println(v13, v5, v3)
            }
            var v18 uint = uint(v15)
            switch (("a" + "\n\t") + "a") {
            case "hello":
                v19 := GenS1{F0: genFn0(1, true)}
                _ = v19
            case "a":
                for fi20 := 0; fi20 < 12; fi20++ {
                    if true {
                        v13 -= -2.25
                    } else {
                        var v21 bool = (127 > -128)
                        fmt.Println(v18)
                        v22 := GenS2{F0: int64(v5)}
                        v23 := GenS3{F0: (v1 - (v1 ^ v1))}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    var v24 uint16 = (v5 - (v5 + v5))
                    var v25 int32 = v1
                    _ = fi20
                    _ = v24
                    _ = v25
                }
            case "foo bar":
                v5 = (v5 ^ 65535)
            }
            _ = fi16
            _ = v18
        }
    }
    v26 := GenS2{F0: 42}
    v27 := GenS4{F0: (v1 + (v1 + (v1 >> ((255) & 63)))), F1: 96}
    var v28 uint = 521
    var v29 string = "foo bar"
    if false {
        v30 := GenS2{F0: 497}
        _ = v30
    }
    if true {
        v13 = 0.0
    } else {
        v2 = v13
    }
    switch 32768 {
    case -269:
        for fi31 := 0; fi31 < 16; fi31++ {
            v29 = (v29 + ((("\n\t" + (("foo bar" + ("a" + "foo bar")) + (("hello" + v29) + "foo bar"))) + "a") + "hello"))
            if true {
                v29 = ""
                fmt.Println(v29, v15, v1)
            } else {
                var v32 uint16 = (^(v5))
                v33 := GenS1{F0: 93}
                v32 += 1
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 16; fi34++ {
                if (0 != (v28 / ((v28) | 1))) {
                    v35 := GenS2{F0: 9223372036854775807}
                    var v36 int8 = -128
                    _ = v35
                    _ = v36
                } else {
                    if false {
                        if true {
                            var v37 uint = 0
                            _ = v37
                        }
                        v1 -= (v1 % (((v1 & -2147483648)) | 1))
                    }
                    fmt.Println(v1, v15)
                    var v38 uint8 = 54
                    _ = v38
                }
                v3 = 1000
                _ = fi34
            }
            _ = fi31
        }
    default:
        v1 += (v1 % ((v1) | 1))
    }
    v3 = v3
    var v39 float64 = 1.5
    var v40 string = ""
    fmt.Println(v40, v3, v29)
    v41 := GenS2{F0: 42}
    v42 := v4[65535]
    fmt.Println(v2, v1, v40)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
