package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int32; F2 string }
type GenS3 struct { F0 uint16; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int8; F1 uint; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 string, p1 int16) string {
    p0 = ("\n\t" + (p0 + (p0 + p0)))
    v7 := GenS1{F0: 4294967295}
    v8 := GenS2{F0: 42, F1: 0, F2: "\n\t"}
    var v9 int8 = -128
    var v10 uint8 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (("a" + "hello") + "hello")
}

func main() {
    var v1 uint = 42
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    var v4 uint32 = 0
    v5 := make(map[string]string)
    v6 := []uint32{(v4 ^ 4294967295), (^(v4))}
    var v11 float64 = float64(0.0)
    if false {
        v12 := GenS3{F0: 255, F1: 42, F2: ""}
        fmt.Println(v4, v4, v1)
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v11, fi13, v2)
            _ = fi13
        }
        _ = v12
    }
    v14 := GenS1{F0: -158}
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS3{F0: 2, F1: 42, F2: "a"}
        _ = fi15
        _ = v16
    }
    var v17 string = (("hello" + ((("" + (("hello" + ("" + "foo bar")) + ("\n\t" + "foo bar"))) + ("hello" + "\n\t")) + ("\n\t" + "a"))) + "a")
    for fi18 := 0; fi18 < 12; fi18++ {
        v11 = v11
        fmt.Println(v4, v11, v4)
        _ = fi18
    }
    if (-9223372036854775808 > 1) {
        fmt.Println(v11, v17)
        switch "hello" {
        case "":
            if len(v6) > 0 {
                v6[0] = (v4 / ((v4) | 1))
            }
        case "a":
            v4 = (v4 << ((v1) & 63))
        default:
            v3[4294967295] = 42
        }
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS3{F0: 42, F1: int64(v1), F2: (((("hello" + v17) + ("\n\t" + (v17 + (v17 + v17)))) + v17) + v17)}
            _ = fi19
            _ = v20
        }
    }
    delete(v5, "hello")
    var v21 uint64 = 2
    var v22 float64 = v11
    fmt.Println(v2, v22, v4)
    var v23 float32 = (1.5 / float32((1e10 * v22)))
    if false {
        var v24 float64 = float64(1e10)
        if (v1 <= (v1 >> ((35) & 63))) {
            v25 := GenS2{F0: -56, F1: (v2 / ((v2) | 1)), F2: v17}
            v26 := GenS4{F0: -23, F1: -679}
            v27 := GenS5{F0: 120, F1: v1, F2: 1}
            _ = v25
            _ = v26
            _ = v27
        }
        v11 = ((((v24 + 3.14159) + v24) / 1.5) * (float64((3.14159 + v23)) - v22))
        _ = v24
    } else {
        var v28 bool = false
        if v28 {
            v29 := GenS6{F0: 83, F1: 824}
            _ = v29
        }
        if (1 < 32767) {
            var v30 uint8 = 1
            _ = v30
        }
        _ = v28
    }
    v31 := GenS3{F0: 0, F1: 42, F2: "\n\t"}
    for fi32 := 0; fi32 < 1; fi32++ {
        if true {
            v33 := GenS4{F0: 42, F1: 42}
            if (!false) {
                var v34 uint8 = 127
                var v35 int32 = int32(v34)
                _ = v34
                _ = v35
            }
            _ = v33
        } else {
            for fi36 := 0; fi36 < 12; fi36++ {
                v5["foo bar"] = "a"
                var v37 int = fi32
                _ = fi36
                _ = v37
            }
            fmt.Println(v4, fi32)
        }
        switch v17 {
        case "hello":
            v11 -= v11
        }
        if len(v6) > 0 {
            v6[0] = uint32(v21)
        }
        _ = fi32
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        for fi39 := 0; fi39 < 8; fi39++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch 229 {
            case -9223372036854775808:
                var v40 float64 = float64((v23 - 1.5))
                _ = v40
            case 0:
                var v41 uint32 = (v4 << (((^(v1))) & 63))
                _ = v41
            }
            _ = fi39
        }
        _ = fi38
    }
    v1 = (v1 - 255)
    v42 := GenS2{F0: 0, F1: 0, F2: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
