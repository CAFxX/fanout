package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "a", "\n\t"}
    var v3 uint8 = 2
    var v4 uint16 = 42
    var v5 float32 = 1e10
    var v6 int64 = 0
    switch "\n\t" {
    case "hello":
        v7 := v1[""]
        _ = v7
    }
    v3 += (v3 * v3)
    v8 := GenS1{F0: -128}
    var v9 uint8 = v3
    if (0.1 >= -2.25) {
        var v10 int8 = -128
        v5 = 1e10
        v9 += v9
        _ = v10
    } else {
        var v11 bool = (127 != 4294967295)
        switch "hello" {
        case "hello":
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(v9, v11, v3)
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 int8 = -11
                    switch (("hello" + "hello") + "hello") {
                    case "a":
                        var v15 float32 = float32(float64(0.1))
                        _ = v15
                    case "":
                        var v16 uint = 42
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS1{F0: 42}
                _ = fi18
                _ = v19
            }
            v11 = v11
            _ = fi17
        }
        _ = v11
    }
    switch (("a" + ("foo bar" + "foo bar")) + "foo bar") {
    case "":
        v2 = append(v2, "a")
    case "hello":
        v4 = 292
    }
    v5 -= 0.0
    for fi20 := 0; fi20 < 5; fi20++ {
        switch -801 {
        case 494:
            v21 := GenS2{F0: 32767}
            _ = v21
        case -2147483648:
            var v22 float32 = v5
            _ = v22
        case -246:
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint16 = (^(v4))
                var v25 int32 = -924
                v5 = v5
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        if (0.1 >= 3.14159) {
            for fi26 := 0; fi26 < 16; fi26++ {
                if (!(!false)) {
                    var v27 uint32 = 4294967295
                    _ = v27
                } else {
                    if (!(true || (42 >= (^(v4))))) {
                        var v28 int64 = 0
                        _ = v28
                    } else {
                        var v29 float64 = -2.25
                        var v30 float32 = v5
                        var v31 uint8 = v3
                        v3 = (v31 >> ((2) & 63))
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                }
                _ = fi26
            }
        }
        _ = fi20
    }
    if false {
        v6 = v6
        v4 = (v4 & v4)
    }
    if true {
        switch ("foo bar" + "\n\t") {
        case "":
            var v32 uint16 = v4
            _ = v32
        case "foo bar":
            switch 4294967295 {
            case 312:
                v1[""] = 2147483648
            default:
                v33 := GenS3{F0: v6}
                _ = v33
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
