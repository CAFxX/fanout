package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{"foo bar", ""}
    var v2 string = "a"
    var v3 bool = (((42 < -216) || (!true)) && (true && (0.1 <= 0.1)))
    var v4 int64 = 780
    var v5 int32 = 1
    v2 = ((v2 + v2) + ("" + ((v2 + (v2 + ("a" + "foo bar"))) + v2)))
    var v6 uint64 = 18446744073709551615
    var v7 uint = 4294967295
    v4 += (v4 + 0)
    fmt.Println(v3, v3, v4)
    if len(v1) > 0 {
        v1[0] = v2
    }
    fmt.Println(v7)
    _ = len(v1)
    switch 308 {
    case 665:
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 string = (("foo bar" + (v2 + v2)) + v2)
            var v10 int64 = (^(v4))
            fmt.Println(v10, v9)
            _ = fi8
            _ = v10
            _ = v9
        }
    case -1:
        fmt.Println(v4, v5, v3)
    case 1:
        _ = len(v1)
    default:
        _ = len(v1)
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    v11 := GenS1{F0: 0, F1: 2, F2: (42 == v6)}
    var v12 float32 = (3.14159 + float32(0.0))
    if (3.14159 > v12) {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint8 = 1
            _ = fi13
            _ = v14
        }
        var v15 uint16 = 65535
        for fi16 := 0; fi16 < 12; fi16++ {
            if true {
                switch fi16 {
                case 463:
                    v12 = v12
                case 32767:
                    v4 += -961
                }
                v3 = false
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int64 = (v4 + v4)
                    var v19 int16 = -1
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                v7 = 433
                switch (v2 + "foo bar") {
                case "\n\t":
                    switch v2 {
                    case "":
                        _ = len(v1)
                    }
                case "hello":
                    v2 = "hello"
                case "foo bar":
                    v20 := GenS2{F0: v4}
                    _ = v20
                }
                v5 = (v5 >> ((65535) & 63))
            }
            v2 = (v2 + "a")
            _ = fi16
        }
        _ = v15
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS1{F0: 44, F1: 4294967295, F2: v3}
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(fi21)
                v2 = ""
                _ = fi23
            }
            v6 = v6
            _ = fi21
            _ = v22
        }
        switch "hello" {
        case "\n\t":
            switch v2 {
            case "a":
                var v24 uint64 = (^(v6))
                _ = v24
            default:
                if (!(!(2 != 128))) {
                    v2 = (v2 + (v2 + (v2 + "a")))
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        default:
            v25 := GenS2{F0: v4}
            _ = v25
        }
        if (v7 <= 1) {
            var v26 int64 = -9223372036854775808
            _ = v26
        } else {
            switch "a" {
            case "":
                v7 += v7
            case "foo bar":
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 bool = v3
                    v2 = ""
                    _ = fi27
                    _ = v28
                }
            default:
                var v29 uint16 = 780
                _ = v29
            }
            v30 := GenS1{F0: 61, F1: 0, F2: true}
            _ = v30
        }
    }
    var v31 int16 = 32767
    fmt.Println(v5, v4, v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
