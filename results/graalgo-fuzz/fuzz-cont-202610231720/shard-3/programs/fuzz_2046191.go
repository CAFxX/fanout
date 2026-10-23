package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0() uint32 {
    if false {
        var v4 uint = 255
        _ = v4
    } else {
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 string = ("\n\t" + "a")
            if false {
                v6 = v6
                var v7 bool = false
                v8 := []int{fi5, (-(fi5)), fi5, (fi5 << ((1) & 63))}
                _ = v7
                _ = v8
            } else {
                var v9 int8 = 1
                if false {
                    var v10 float32 = 3.14159
                    v6 = (v6 + ((((v6 + (v6 + v6)) + v6) + "\n\t") + ("foo bar" + "foo bar")))
                    _ = v10
                }
                v9 = (v9 << ((352) & 63))
                _ = v9
            }
            _ = fi5
            _ = v6
        }
        var v11 int16 = -32768
        switch "foo bar" {
        case "":
            fmt.Println(v11)
        case "a":
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(v11, v11)
                v13 := GenS1{F0: 0, F1: -1, F2: 143}
                _ = fi12
                _ = v13
            }
        case "hello":
            switch 1 {
            case 345:
                v11 -= v11
            case 4294967295:
                var v14 float64 = 3.14159
                _ = v14
            }
        }
        _ = v11
    }
    var v15 uint8 = 2
    if false {
        var v16 string = ""
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := []int{(-(fi17)), 32768, (fi17 >> ((2) & 63))}
            if false {
                v20 := GenS1{F0: -1, F1: 2147483647, F2: 4294967295}
                if ((304 != 1) || (!false)) {
                    v21 := GenS1{F0: 42, F1: -1, F2: 593}
                    v22 := GenS1{F0: 0, F1: 386, F2: 42}
                    v15 += (v15 - v15)
                    _ = v21
                    _ = v22
                }
                _ = v20
            } else {
                v15 = (v15 << ((255) & 63))
                v15 = 244
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v23 int64 = -9223372036854775808
    _ = v15
    _ = v23
    return 603
}
func genFn1() int32 {
    if (2147483648 != 2) {
        var v24 string = "foo bar"
        _ = v24
    }
    switch -899 {
    case 159:
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 int = -1
            v27 := GenS1{F0: 9223372036854775807, F1: -376, F2: 925}
            var v28 uint64 = 1
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
    case 32768:
        var v29 uint8 = 2
        _ = v29
    case -9223372036854775808:
        var v30 float64 = 0.0
        _ = v30
    default:
        v31 := GenS2{F0: 42, F1: 32767}
        _ = v31
    }
    v32 := GenS2{F0: 127, F1: 927}
    v33 := GenS1{F0: -534, F1: 1, F2: 4294967295}
    _ = v32
    _ = v33
    return 2147483647
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = 1
    v3 := make(map[uint32]bool)
    v34 := GenS1{F0: 0, F1: (v1 / ((v1) | 1)), F2: genFn0()}
    var v35 float32 = 0.0
    var v36 uint8 = 128
    v1 = v1
    v37 := GenS1{F0: 9223372036854775807, F1: -1, F2: 671}
    if false {
        var v38 int = 128
        _ = v38
    } else {
        fmt.Println(v2, v36, v2)
    }
    fmt.Println(v36, v36)
    fmt.Println(v1)
    v39 := GenS1{F0: 9223372036854775807, F1: 507, F2: 1}
    v40 := GenS2{F0: 1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v40
}
