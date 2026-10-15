package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 int32; F1 int32; F2 int8 }
func genFn0(p0 uint16) uint32 {
    var v4 uint16 = 382
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint16 = v4
        var v7 uint64 = 18446744073709551615
        v4 = v4
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 string = ("hello" + ("\n\t" + ""))
    if (991 <= 127) {
        var v9 int8 = 40
        if false {
            for fi10 := 0; fi10 < 7; fi10++ {
                if (0.1 < 3.14159) {
                    var v11 int8 = -1
                    v12 := GenS1{F0: 2}
                    if false {
                        var v13 int8 = v9
                        v14 := []int{(fi10 / ((2147483647) | 1)), -974, (fi10 ^ (fi10 | fi10))}
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
                v15 := GenS2{F0: (0 < 4294967295)}
                fmt.Println(fi10)
                _ = fi10
                _ = v15
            }
        }
        var v16 uint = 42
        _ = v16
        _ = v9
    }
    v17 := GenS2{F0: true}
    _ = p0
    _ = v17
    _ = v4
    _ = v8
    return 2
}

func main() {
    v1 := []uint32{2, 977, 2147483647}
    v2 := make(map[uint32]string)
    v3 := []string{"\n\t", "\n\t", "hello"}
    for fi18 := 0; fi18 < 2; fi18++ {
        if false {
            var v19 int8 = 127
            _ = v19
        }
        fmt.Println(fi18)
        _ = fi18
    }
    var v20 uint = 2
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            v23 := GenS1{F0: 657}
            if (false && (true || false)) {
                v20 += 0
                v2[genFn0(284)] = ""
            }
            _ = fi22
            _ = v23
        }
        switch fi21 {
        case -1:
            var v24 int16 = -32768
            _ = v24
        default:
            var v25 float32 = ((float32((0.1 + 3.14159)) + 3.14159) / (((0.1 / (3.14159 - 0.1)) - 0.0) / (1.5 / 1e10)))
            _ = v25
        }
        _ = fi21
    }
    fmt.Println(v20)
    if (!("a" != ("a" + (("a" + "hello") + "")))) {
        switch ("" + "hello") {
        case "hello":
            v20 = v20
        case "foo bar":
            var v26 bool = false
            _ = v26
        case "":
            switch "hello" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        default:
            var v27 int16 = -32768
            _ = v27
        }
        v2[739] = ""
        v28 := GenS2{F0: (4294967295 <= 0)}
        _ = v28
    } else {
        switch -2147483648 {
        case 1:
            v29 := GenS3{F0: -128, F1: 32767, F2: 0}
            _ = v29
        case -323:
            switch "a" {
            case "foo bar":
                var v30 int8 = -1
                _ = v30
            case "a":
                v20 = v20
            case "hello":
                v20 = 1
            default:
                v20 = 1
            }
        case 0:
            var v31 int8 = 42
            _ = v31
        }
        if true {
            v20 = (v20 + (v20 + (^(v20))))
        }
        for fi32 := 0; fi32 < 13; fi32++ {
            for fi33 := 0; fi33 < 2; fi33++ {
                switch "" {
                case "hello":
                    var v34 uint64 = 42
                    _ = v34
                case "foo bar":
                    var v35 int16 = -32768
                    _ = v35
                }
                if true {
                    if false {
                        v2[2147483647] = ("" + "foo bar")
                        v36 := GenS3{F0: -101, F1: 530, F2: 128}
                        v37 := GenS2{F0: (!(true && (true || (!(!true)))))}
                        _ = v36
                        _ = v37
                    }
                    var v38 uint16 = 1
                    _ = v38
                }
                _ = fi33
            }
            _ = fi32
        }
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        switch -1 {
        case 127:
            v20 = (v20 | (v20 << (((v20 & (v20 >> ((v20) & 63)))) & 63)))
        case -53:
            fmt.Println(v20)
        case 42:
            v20 -= v20
        default:
            for fi40 := 0; fi40 < 16; fi40++ {
                v41 := GenS3{F0: 100, F1: fi39, F2: 195}
                var v42 uint = (v20 << (((v20 * (v20 ^ (v20 >> (((v20 | (v20 / ((v20) | 1)))) & 63))))) & 63))
                switch -2147483648 {
                case 42:
                    v42 -= (v42 << ((4294967295) & 63))
                default:
                    if (!true) {
                        var v43 float32 = 1e10
                        var v44 uint8 = 201
                        v44 += v44
                        _ = v43
                        _ = v44
                    } else {
                        v20 -= (v42 + v42)
                        var v45 uint32 = 2
                        v45 = (v45 - v45)
                        _ = v45
                    }
                }
                v46 := GenS2{F0: (false || ((238 > -1) || (-12 != 9223372036854775807)))}
                _ = fi40
                _ = v41
                _ = v42
                _ = v46
            }
        }
        _ = fi39
    }
    for fi47 := 0; fi47 < 5; fi47++ {
        v48 := GenS3{F0: -1, F1: (-(fi47)), F2: 127}
        _ = fi47
        _ = v48
    }
    v20 = v20
    for fi49 := 0; fi49 < 9; fi49++ {
        var v50 int32 = 42
        _ = fi49
        _ = v50
    }
    fmt.Println(v20)
    var v51 uint16 = 255
    v52 := GenS4{F0: -2147483648, F1: 1, F2: 127}
    v51 = (v51 >> (((v20 >> ((v20) & 63))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v51
    _ = v52
}
