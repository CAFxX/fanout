package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 uint; F2 int }

func main() {
    var v1 uint32 = 2
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    if true {
        var v5 string = "a"
        _ = v5
    } else {
        v2 += v2
        _ = len(v3)
        v6 := GenS1{F0: true, F1: true}
        _ = v6
    }
    v7 := GenS1{F0: false, F1: true}
    if true {
        var v8 int = 2147483647
        var v9 uint64 = 0
        var v10 int = -2147483648
        _ = v10
        _ = v8
        _ = v9
    } else {
        if false {
            v1 = 800
        }
    }
    v1 = (v1 + 4294967295)
    v11 := GenS2{F0: 1}
    v12 := []int{-2147483648}
    v13 := GenS3{F0: 1, F1: 255, F2: 711}
    v14 := GenS4{F0: (("foo bar" + "hello") + "\n\t"), F1: ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "a")))}
    if true {
        v1 -= (v1 ^ 65535)
    } else {
        v15 := GenS2{F0: 0}
        _ = v15
    }
    _ = len(v4)
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    for fi16 := 0; fi16 < 13; fi16++ {
        switch (fi16 % (((fi16 | (fi16 >> ((255) & 63)))) | 1)) {
        case -2147483648:
            switch ("" + "hello") {
            case "":
                if ("" >= ((("foo bar" + ("a" + (("foo bar" + "foo bar") + "foo bar"))) + "") + "\n\t")) {
                    if true {
                        v1 -= 480
                    } else {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v2 -= v2
                            _ = fi17
                        }
                        fmt.Println(v1)
                        switch 42 {
                        case -25:
                            var v18 float32 = 1e10
                            _ = v18
                        default:
                            var v19 uint64 = 325
                            _ = v19
                        }
                    }
                } else {
                    var v20 uint16 = 0
                    var v21 uint32 = v1
                    var v22 int16 = -487
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
        case 716:
            delete(v4, "hello")
        case 2147483647:
            var v23 uint = 1
            _ = v23
        default:
            switch "hello" {
            case "foo bar":
                var v24 int16 = -789
                _ = v24
            default:
                v25 := GenS5{F0: 255, F1: 1, F2: fi16}
                _ = v25
            }
        }
        _ = fi16
    }
    v12 = append(v12, 32767)
    for fi26 := 0; fi26 < 19; fi26++ {
        delete(v3, v1)
        v1 = (v1 << ((461) & 63))
        var v27 uint32 = (v1 ^ (v1 * (v1 + (v1 / ((v1) | 1)))))
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
