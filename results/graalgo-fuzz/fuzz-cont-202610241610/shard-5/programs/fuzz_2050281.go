package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
func genFn0() string {
    if (1e10 != (0.1 + float32((1.5 + ((1.5 * ((1e10 / 1.5) / 3.14159)) / 1.0))))) {
        switch ("hello" + "\n\t") {
        case "foo bar":
            v7 := GenS1{F0: -848}
            _ = v7
        case "\n\t":
            var v8 int8 = -1
            _ = v8
        case "hello":
            v9 := []int{9223372036854775807, 32768}
            _ = v9
        }
    }
    v10 := []int{-272, 128, 32767}
    _ = v10
    return ("\n\t" + "a")
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint16 = 474
    var v3 int8 = 127
    var v4 uint8 = 127
    v5 := make(map[string]uint32)
    var v6 uint = 827
    v11 := GenS2{F0: (v4 & v4)}
    var v12 int = int(v4)
    var v13 int64 = 0
    v6 = 4294967295
    if (v13 <= -1) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v12 -= v12
    } else {
        v2 = 29
        v14 := GenS3{F0: ""}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 float32 = 3.14159
        if (!true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v12 += (fi15 >> ((uint(v4)) & 63))
            var v17 int64 = (v13 << ((v6) & 63))
            _ = v17
        } else {
            switch (fi15 - (v12 + 0)) {
            case -1:
                if false {
                    v18 := GenS1{F0: fi15}
                    v4 = uint8(v6)
                    v13 += 42
                    _ = v18
                } else {
                    var v19 int32 = 195
                    var v20 uint8 = 191
                    var v21 bool = true
                    _ = v19
                    _ = v20
                    _ = v21
                }
            case 32768:
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 int64 = int64(v4)
                    _ = len(v1)
                    _ = fi22
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 int = v12
                    v13 = v13
                    var v26 float64 = float64((v16 / v16))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            var v27 uint16 = (v2 & (v2 + uint16(v4)))
            _ = v27
        }
        _ = fi15
        _ = v16
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = (^(v3))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 = 42
    var v28 string = "a"
    var v29 int32 = 42
    fmt.Println(v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
