package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]uint32)
    var v3 string = ("a" + "a")
    var v4 uint16 = 65535
    var v5 int = -1
    v6 := []string{v3, v3}
    if len(v6) > 0 {
        v6[0] = ("hello" + (v3 + v3))
    }
    fmt.Println(v5, v3, v5)
    v3 = "a"
    _ = len(v6)
    v3 = (("hello" + v3) + ((v3 + (("" + v3) + "hello")) + ""))
    v7 := GenS1{F0: 65535}
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v8 float64 = 3.14159
        _ = v8
    } else {
        if (false && false) {
            v5 = v5
        }
    }
    v5 = 4294967295
    switch (-(v5)) {
    case -1:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case -9223372036854775808:
        delete(v2, 627)
    }
    v5 = v5
    for fi9 := 0; fi9 < 13; fi9++ {
        switch v5 {
        case -131:
            var v10 uint32 = uint32(v4)
            _ = v10
        case 95:
            switch ("" + ((((v3 + v3) + v3) + v3) + v3)) {
            case "foo bar":
                v3 = "hello"
            case "hello":
                v1 -= v1
            case "a":
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 float32 = (1.5 / 1e10)
                    v2[42] = uint32(v4)
                    v12 += -2.25
                    fmt.Println(fi9, v12, fi11)
                    var v13 int8 = -1
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                if true {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 uint32 = 2
                        v4 = (v4 | v4)
                        _ = fi14
                        _ = v15
                    }
                    if true {
                        var v16 uint32 = 2
                        var v17 string = v3
                        for fi18 := 0; fi18 < 15; fi18++ {
                            var v19 int8 = 0
                            switch fi9 {
                            case -1:
                                var v20 uint32 = (^(v16))
                                _ = v20
                            case 0:
                                var v21 int8 = (v19 | (v19 & 88))
                                _ = v21
                            case 127:
                                delete(v2, v16)
                            }
                            _ = fi18
                            _ = v19
                        }
                        _ = v16
                        _ = v17
                    } else {
                        var v22 int = -959
                        var v23 uint8 = 2
                        var v24 int32 = -462
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS1{F0: 799}
                    _ = v25
                } else {
                    if (((true || (0.0 != 0.1)) && false) || ((9223372036854775807 != 719) && true)) {
                        v26 := GenS1{F0: (^(v4))}
                        v5 += v5
                        var v27 float64 = 1e10
                        _ = v26
                        _ = v27
                    }
                    v4 += (v4 - v4)
                }
            }
        }
        if true {
            v3 = "\n\t"
            for fi28 := 0; fi28 < 4; fi28++ {
                v3 = ""
                var v29 float64 = -2.25
                switch (fi9 ^ v5) {
                case 9223372036854775807:
                    var v30 uint64 = 18446744073709551615
                    v3 = ("foo bar" + v3)
                    _ = v30
                case -554:
                    v31 := v2[325]
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
            var v32 int64 = -967
            _ = v32
        } else {
            var v33 uint64 = 0
            var v34 int = (fi9 / (((fi9 % ((v5) | 1))) | 1))
            _ = v33
            _ = v34
        }
        v35 := GenS1{F0: v4}
        _ = fi9
        _ = v35
    }
    v4 = v4
    var v36 int32 = 2147483647
    var v37 int = int(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
