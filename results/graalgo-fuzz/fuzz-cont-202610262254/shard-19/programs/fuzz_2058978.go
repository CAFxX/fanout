package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int8) uint {
    p0 += (p0 ^ p0)
    var v5 uint16 = 65535
    switch "hello" {
    case "":
        var v6 float32 = (float32(0.0) / (0.1 * 0.1))
        _ = v6
    case "hello":
        var v7 string = ("\n\t" + "foo bar")
        _ = v7
    case "a":
        switch -2147483648 {
        case -392:
            v8 := GenS1{F0: 0}
            _ = v8
        case 42:
            fmt.Println(p0, p0)
        case 959:
            var v9 uint = 1
            _ = v9
        default:
            v5 -= (^(v5))
        }
    default:
        p0 += -1
    }
    var v10 string = "hello"
    v10 = ("hello" + ("hello" + (v10 + (v10 + ((v10 + v10) + v10)))))
    _ = p0
    _ = v10
    _ = v5
    return 900
}
func genFn1() int16 {
    var v11 string = ""
    if true {
        if true {
            if true {
                if false {
                    var v12 float32 = ((0.1 - (1e10 + 0.0)) * ((0.1 / float32((0.1 - (1.5 / (-2.25 + 1.5))))) + 3.14159))
                    switch v11 {
                    case "a":
                        v13 := GenS1{F0: 911}
                        _ = v13
                    }
                    fmt.Println(v11, v11)
                    _ = v12
                } else {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v11 = ("\n\t" + (v11 + v11))
                        _ = fi14
                    }
                    var v15 uint = 65535
                    _ = v15
                }
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 float64 = -2.25
                    var v18 int8 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            v19 := []int{-1, 32767}
            _ = v19
        } else {
            if (!false) {
                v20 := GenS2{F0: 1}
                _ = v20
            } else {
                var v21 uint16 = 2
                v22 := GenS1{F0: 0}
                var v23 float64 = (float64(0.0) * 0.0)
                _ = v21
                _ = v22
                _ = v23
            }
            var v24 int64 = -827
            _ = v24
        }
    } else {
        switch v11 {
        case "":
            v11 = ("\n\t" + ("foo bar" + ""))
        case "\n\t":
            var v25 uint32 = 854
            _ = v25
        default:
            fmt.Println(v11)
        }
    }
    _ = v11
    return 198
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = (!(!(1 <= 255)))
    v3 := make(map[uint32]int)
    var v4 uint = 42
    var v26 int32 = -2147483648
    var v27 float32 = (1e10 / 1e10)
    v28 := GenS2{F0: 4294967295}
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 int64 = 9223372036854775807
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 int64 = (v30 >> (((v4 ^ v4)) & 63))
            _ = fi31
            _ = v32
        }
        v33 := []int{(fi29 ^ 42), 2147483647, (fi29 ^ 32768), 128}
        _ = fi29
        _ = v30
        _ = v33
    }
    var v34 bool = (127 == 97)
    if (196 < v4) {
        v3[2147483648] = 32767
    }
    fmt.Println(v34, v26, v27)
    v35 := GenS2{F0: 4294967295}
    fmt.Println(v4, v26)
    switch "foo bar" {
    case "foo bar":
        v3[1] = 1
    case "":
        v36 := []int{342}
        _ = v36
    case "hello":
        var v37 float64 = 0.1
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := GenS2{F0: 65535}
        v40 := GenS3{F0: 2, F1: 127}
        v26 += (v26 - v26)
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println(v4)
    switch ("" + ("hello" + ("hello" + ""))) {
    case "foo bar":
        v1 -= v1
    }
    fmt.Println(v2, v2, v34)
    for fi41 := 0; fi41 < 2; fi41++ {
        v1 = v1
        v3[2147483647] = int(v1)
        var v42 uint16 = 793
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        for fi44 := 0; fi44 < 10; fi44++ {
            switch fi43 {
            case 4294967295:
                v45 := []int{}
                _ = v45
            case 2147483647:
                v46 := v3[65535]
                _ = v46
            case 1:
                switch int(v26) {
                case 42:
                    v47 := GenS2{F0: 4294967295}
                    _ = v47
                case 127:
                    v34 = v34
                }
            }
            v48 := GenS2{F0: 0}
            _ = len(v3)
            _ = fi44
            _ = v48
        }
        v49 := GenS3{F0: 343, F1: 255}
        _ = fi43
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
