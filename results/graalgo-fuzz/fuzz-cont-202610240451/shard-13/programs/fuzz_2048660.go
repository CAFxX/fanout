package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]string)
    var v3 string = ""
    v4 := []string{v3, v3, ("a" + "a"), (((v3 + ("hello" + v3)) + v3) + v3)}
    v5 := make(map[int]int)
    var v6 int8 = 0
    v1 = v1
    switch -2147483648 {
    case 32767:
        if (0 >= 721) {
            v3 = v3
            for fi7 := 0; fi7 < 14; fi7++ {
                if false {
                    v5[(fi7 + fi7)] = (fi7 ^ 42)
                    var v8 uint = 255
                    _ = v8
                } else {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            var v9 int8 = v6
                            _ = v9
                        default:
                            v10 := GenS1{F0: v3}
                            _ = v10
                        }
                    }
                    var v11 int = (fi7 - 42)
                    _ = v11
                }
                var v12 int32 = -2147483648
                _ = fi7
                _ = v12
            }
        } else {
            v5[32768] = 128
            if len(v4) > 0 {
                v4[0] = (v3 + v3)
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 float64 = 1.5
                _ = fi13
                _ = v14
            }
        }
    case -39:
        delete(v5, -2147483648)
    case -2147483648:
        _ = len(v2)
    default:
        v15 := GenS1{F0: v3}
        _ = v15
    }
    v1 = 127
    if false {
        switch (v3 + "hello") {
        case "hello":
            var v16 int8 = v6
            _ = v16
        }
        if (0.0 == ((0.1 + 3.14159) - 3.14159)) {
            fmt.Println(v1, v3, v1)
            v5[0] = 569
        } else {
            var v17 uint32 = 14
            _ = v17
        }
    }
    fmt.Println(v3, v3, v1)
    v6 -= (v6 >> ((4294967295) & 63))
    if len(v4) > 0 {
        v4[0] = (v3 + "\n\t")
    }
    if (!(true && (((int16(v6) != -1) || ((true || (1 <= -1)) || (127 > v1))) || true))) {
        if true {
            switch v3 {
            case "\n\t":
                var v18 int32 = 1
                _ = v18
            }
            for fi19 := 0; fi19 < 16; fi19++ {
                if (14 == 660) {
                    var v20 uint64 = 42
                    _ = v20
                } else {
                    v21 := GenS2{F0: -32768}
                    fmt.Println(fi19)
                    _ = v21
                }
                _ = fi19
            }
            v3 = v3
        } else {
            v22 := GenS2{F0: -1}
            v1 = v1
            v6 = 79
            _ = v22
        }
        v5[32768] = -2147483648
    } else {
        var v23 string = v3
        if false {
            var v24 int = 2147483647
            v1 = v1
            var v25 bool = true
            _ = v24
            _ = v25
        }
        v26 := GenS1{F0: v23}
        _ = v23
        _ = v26
    }
    var v27 float64 = float64(-2.25)
    v28 := GenS1{F0: v3}
    var v29 float32 = 0.1
    if true {
        v1 += 0
    } else {
        fmt.Println(v1)
    }
    v6 = (v6 + (v6 << ((4294967295) & 63)))
    v27 -= v27
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
