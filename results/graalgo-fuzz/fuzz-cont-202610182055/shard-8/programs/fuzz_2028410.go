package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 string; F1 string; F2 uint8 }
func genFn0(p0 bool) bool {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = p0
        _ = fi7
    }
    var v8 float64 = (((0.0 - 0.0) * (((1e10 + (1.5 + -2.25)) - 3.14159) + 3.14159)) / 0.1)
    v8 -= (v8 + (v8 / 1.0))
    _ = p0
    _ = v8
    return true
}
func genFn1(p0 uint, p1 int, p2 int8) int {
    fmt.Println(p2)
    p0 = (p0 + 42)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := []uint32{4294967295, 2147483648}
    v2 := []int{128, -217, 127, 1}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 128
    var v5 bool = false
    var v6 string = ("hello" + "hello")
    v6 = (v6 + v6)
    var v9 int32 = -2147483648
    if v5 {
        switch ("\n\t" + "hello") {
        case "\n\t":
            fmt.Println(v6)
        case "hello":
            for fi10 := 0; fi10 < 12; fi10++ {
                for fi11 := 0; fi11 < 7; fi11++ {
                    if true {
                        var v12 bool = v5
                        v13 := GenS1{F0: 4294967295}
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(v6)
                    _ = fi11
                }
                v6 = v6
                _ = fi10
            }
        case "":
            v14 := GenS1{F0: 0}
            _ = v14
        default:
            var v15 int32 = -168
            _ = v15
        }
        switch 42 {
        case 2147483647:
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 int = fi16
                v9 += (v9 & v9)
                _ = fi16
                _ = v17
            }
        default:
            for fi18 := 0; fi18 < 20; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    v20 := GenS1{F0: 647}
                    var v21 int8 = 42
                    var v22 int16 = 1000
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v9 = (v9 << ((65535) & 63))
                _ = fi18
            }
        }
        v4 -= v4
    }
    var v23 uint = 2
    for fi24 := 0; fi24 < 5; fi24++ {
        v4 = (v4 | uint8(v23))
        v4 -= (^(v4))
        _ = fi24
    }
    v3[42] = (!false)
    var v25 uint8 = (v4 % ((0) | 1))
    if ((v25 % ((128) | 1)) < (v4 & (v25 & v25))) {
        switch ("\n\t" + "hello") {
        case "hello":
            v9 = v9
        case "\n\t":
            switch ("" + v6) {
            case "a":
                switch 386 {
                case 128:
                    if ((v25 ^ (v4 / (((v25 - v4)) | 1))) == (v4 % (((v4 * v4)) | 1))) {
                        v6 = (v6 + "")
                        v26 := GenS1{F0: 498}
                        _ = v26
                    } else {
                        if false {
                            fmt.Println(v25, v25)
                        } else {
                            fmt.Println(v23)
                            v4 -= v4
                            v6 = (v6 + v6)
                        }
                    }
                case 298:
                    if (((3.14159 - (((0.1 * 1e10) - 1.5) - ((-2.25 * 1.5) - 1.5))) > 1e10) && (1000 == -828)) {
                        v6 = ("hello" + (v6 + v6))
                        switch int(v23) {
                        case 9223372036854775807:
                            var v27 string = v6
                            _ = v27
                        case -204:
                            v9 = -1
                        default:
                            v6 = (((v6 + (v6 + ((v6 + v6) + v6))) + (v6 + v6)) + ((v6 + (v6 + (v6 + v6))) + ("a" + "hello")))
                        }
                    }
                default:
                    var v28 int32 = -2147483648
                    _ = v28
                }
            case "foo bar":
                v29 := GenS2{F0: 792}
                _ = v29
            }
        case "foo bar":
            switch 32767 {
            case 32768:
                if (((42 < 0) && v5) && false) {
                    var v30 uint8 = v4
                    switch 32768 {
                    case 2147483647:
                        fmt.Println(v6)
                    case -726:
                        var v31 int8 = 100
                        _ = v31
                    }
                    _ = v30
                }
            case -9223372036854775808:
                v4 -= 0
            }
        default:
            v23 = (v23 - 65535)
        }
    }
    v32 := GenS3{F0: 488, F1: true}
    if v5 {
        if ((-1 != -62) || (v5 || false)) {
            var v33 int8 = 127
            v23 += (v23 << ((v23) & 63))
            _ = v33
        }
        if (1 >= 2147483648) {
            var v34 float32 = 1.5
            var v35 string = (v6 + ((v6 + ("a" + "")) + "foo bar"))
            _ = v34
            _ = v35
        }
    }
    v9 = v9
    switch ("\n\t" + (((v6 + v6) + v6) + ("a" + (v6 + v6)))) {
    case "foo bar":
        v9 = v9
    default:
        for fi36 := 0; fi36 < 15; fi36++ {
            v25 -= (v25 * v4)
            _ = fi36
        }
    }
    switch 127 {
    case 755:
        var v37 uint16 = 0
        _ = v37
    case 0:
        v38 := GenS2{F0: 0}
        _ = v38
    case 1:
        v3[702] = (v23 >= 255)
    default:
        if (!(2 < 4294967295)) {
            var v39 int16 = 0
            v40 := GenS4{F0: ("\n\t" + "foo bar"), F1: v6, F2: v25}
            var v41 int64 = 1
            _ = v39
            _ = v40
            _ = v41
        }
    }
    v42 := GenS4{F0: "", F1: "foo bar", F2: v4}
    switch v6 {
    case "a":
        if (!(!true)) {
            for fi43 := 0; fi43 < 18; fi43++ {
                v9 += 207
                fmt.Println(v4)
                _ = fi43
            }
            v4 = (v25 >> ((567) & 63))
            v44 := GenS4{F0: (v6 + (v6 + ((((v6 + v6) + (v6 + v6)) + v6) + "foo bar"))), F1: ((((v6 + ("" + v6)) + "foo bar") + "\n\t") + ("" + v6)), F2: 42}
            _ = v44
        }
    case "\n\t":
        var v45 float64 = 1.5
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
