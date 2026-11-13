package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 uint16, p1 uint8) string {
    switch 32768 {
    case 344:
        var v5 int32 = 42
        _ = v5
    case 32767:
        var v6 string = ("foo bar" + "")
        _ = v6
    case -627:
        if false {
            fmt.Println(p0, p1)
            if ((("" + "hello") + ("hello" + "")) <= ("\n\t" + "hello")) {
                fmt.Println(p0, p1)
            } else {
                p1 -= p1
                fmt.Println(p0, p1)
                var v7 uint = 4294967295
                _ = v7
            }
        } else {
            var v8 int16 = -307
            v9 := GenS1{F0: ""}
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint = 1
                var v12 float32 = (((((float32((1e10 * 0.0)) * 0.1) - ((0.0 + 0.0) + (-2.25 * 0.1))) - 1e10) * 1.5) * (((float32(((float64(1e10) * 0.0) - 0.1)) / 1.5) / 1.5) + ((1.5 - float32((1e10 + 1.5))) * 0.0)))
                if true {
                    var v13 uint32 = 1
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        }
    default:
        if (!(1 < 42)) {
            v14 := GenS2{F0: (p1 | (p1 | 0)), F1: -863, F2: 0}
            p1 += p1
            var v15 uint = 0
            _ = v14
            _ = v15
        }
    }
    var v16 int16 = 1000
    v17 := []int{32767}
    v18 := GenS2{F0: p1, F1: (v16 >> ((2) & 63)), F2: 87}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    v3 := []uint32{42}
    v4 := make(map[uint32]bool)
    v2 = v2
    for fi19 := 0; fi19 < 9; fi19++ {
        switch int(v2) {
        case 9223372036854775807:
            switch "foo bar" {
            case "":
                v20 := v4[65535]
                _ = v20
            case "\n\t":
                v2 -= v2
            case "a":
                v2 = (v2 | v2)
            }
        default:
            if true {
                var v21 int8 = 1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v21
            }
        }
        _ = fi19
    }
    _ = len(v1)
    v1[183] = ""
    switch 9223372036854775807 {
    case -609:
        var v22 int8 = 127
        _ = v22
    }
    var v23 uint8 = (v2 - (v2 + (v2 - (v2 ^ v2))))
    v2 -= 0
    if true {
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v2)
            _ = fi24
        }
        switch -2147483648 {
        case 128:
            var v25 uint64 = 4294967295
            _ = v25
        case 9223372036854775807:
            v2 = 71
        case 4294967295:
            v26 := GenS1{F0: (("" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t"))}
            _ = v26
        default:
            for fi27 := 0; fi27 < 6; fi27++ {
                fmt.Println(v23, fi27, fi27)
                _ = fi27
            }
        }
    } else {
        v3 = append(v3, 0)
        for fi28 := 0; fi28 < 3; fi28++ {
            v29 := GenS1{F0: ("\n\t" + "a")}
            var v30 uint64 = 18446744073709551615
            _ = fi28
            _ = v29
            _ = v30
        }
        if (false || false) {
            var v31 float32 = 3.14159
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint16 = 147
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
