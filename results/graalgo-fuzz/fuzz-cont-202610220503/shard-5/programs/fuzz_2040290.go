package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint8 = 0
    v2 := []string{"\n\t"}
    var v3 uint = 255
    v4 := make(map[string]bool)
    if true {
        v5 := GenS1{F0: 820, F1: 310}
        var v6 int16 = -32768
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: uint32(v3), F1: 120}
            v9 := GenS2{F0: false, F1: 1000}
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(v3)
        switch 0 {
        case 2147483647:
            var v10 bool = (0 == 994)
            _ = v10
        default:
            if true {
                fmt.Println(v3, v3)
            } else {
                v11 := GenS2{F0: (!true), F1: 42}
                var v12 int16 = 32767
                _ = v11
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 float32 = (-2.25 * 1e10)
        _ = len(v2)
        if false {
            var v15 int32 = 42
            var v16 uint16 = 65535
            _ = v15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v3 -= 324
    v17 := GenS3{F0: 264}
    for fi18 := 0; fi18 < 18; fi18++ {
        if ((949 <= 1) && (true || (true && true))) {
            delete(v4, ("hello" + "hello"))
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 int16 = -1
                if (false && false) {
                    var v21 float32 = float32(3.14159)
                    v3 = v3
                    if true {
                        switch "foo bar" {
                        case "\n\t":
                            var v22 uint64 = 4294967295
                            _ = v22
                        case "hello":
                            v23 := GenS4{F0: -128}
                            _ = v23
                        }
                        v21 -= v21
                    }
                    var v24 bool = (4294967295 <= 0)
                    _ = v21
                    _ = v24
                } else {
                    var v25 uint16 = 464
                    _ = len(v4)
                    var v26 float32 = (-2.25 * ((1e10 / 0.1) + float32((float64(0.1) * 3.14159))))
                    _ = v25
                    _ = v26
                }
                if (false || true) {
                    var v27 int64 = 42
                    var v28 uint = (v3 >> ((v3) & 63))
                    _ = v27
                    _ = v28
                } else {
                    _ = len(v2)
                    var v29 float64 = 1.5
                    _ = v29
                }
                _ = fi19
                _ = v20
            }
        } else {
            v30 := v4[("a" + "foo bar")]
            var v31 uint64 = 315
            v30 = false
            _ = v30
            _ = v31
        }
        v32 := v4[("" + ("foo bar" + "hello"))]
        _ = fi18
        _ = v32
    }
    v1 = v1
    var v33 string = ""
    v34 := v4[(v33 + "foo bar")]
    v35 := GenS1{F0: 2, F1: 42}
    fmt.Println(v33)
    var v36 int32 = 431
    for fi37 := 0; fi37 < 2; fi37++ {
        switch (("a" + v33) + v33) {
        case "\n\t":
            v33 = (v33 + v33)
        }
        _ = fi37
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v38 float32 = -2.25
    switch v33 {
    case "\n\t":
        var v39 uint8 = v1
        _ = v39
    }
    fmt.Println(v36)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v4
}
