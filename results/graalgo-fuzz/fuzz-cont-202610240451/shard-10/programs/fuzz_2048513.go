package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8) int64 {
    p0 -= 42
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 bool = true
        if v6 {
            var v7 float32 = float32(((3.14159 / -2.25) - (float64(0.1) / (1e10 + 0.0))))
            _ = v7
        }
        var v8 uint32 = 1
        _ = fi5
        _ = v6
        _ = v8
    }
    _ = p0
    return -1
}

func main() {
    v1 := []uint32{571}
    v2 := make(map[string]bool)
    v3 := []string{("\n\t" + ""), "", "hello"}
    v4 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float64 = 0.0
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        switch "a" {
        case "foo bar":
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                switch (fi12 ^ fi12) {
                case -836:
                    delete(v2, ("hello" + "foo bar"))
                case 42:
                    fmt.Println(fi12)
                }
            case "":
                if false {
                    var v13 int32 = -2147483648
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v4[1] = (!(!true))
                        v15 := GenS1{F0: 2, F1: 128}
                        var v16 uint = 255
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v17 float32 = 1.5
                    _ = v13
                    _ = v17
                }
            case "\n\t":
                var v18 float32 = ((-2.25 * -2.25) / -2.25)
                _ = v18
            default:
                fmt.Println(fi12)
            }
        case "hello":
            var v19 int32 = -2147483648
            _ = v19
        case "":
            var v20 int8 = -128
            _ = v20
        }
        _ = fi12
    }
    var v21 bool = (false && (106 < -1))
    var v22 string = ("foo bar" + "hello")
    var v23 uint = 1
    switch "hello" {
    case "a":
        var v24 int32 = 2147483647
        _ = v24
    case "\n\t":
        v25 := GenS1{F0: 0, F1: 166}
        _ = v25
    case "foo bar":
        v3 = append(v3, "hello")
    }
    v26 := GenS2{F0: 218}
    var v27 uint = v23
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 uint = v27
        _ = fi28
        _ = v29
    }
    var v30 bool = v21
    v22 = ("foo bar" + v22)
    v23 = v23
    fmt.Println(v21, v21, v21)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
}
