package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{}
    var v3 float64 = 0.1
    v4 := []string{}
    v3 = 3.14159
    var v5 uint = 255
    v3 = v3
    v6 := GenS1{F0: ""}
    v1 = (^(v1))
    v3 += ((v3 + v3) + v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 -= (v5 & 4294967295)
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            if len(v4) > 0 {
                v4[0] = "a"
            }
        } else {
            fmt.Println(v1)
            var v8 int8 = 42
            _ = v8
        }
        v9 := GenS2{F0: 1}
        _ = fi7
        _ = v9
    }
    if (!(float32((0.0 * 1.5)) <= 3.14159)) {
        if ((true && true) || (!false)) {
            v10 := GenS3{F0: -1}
            if false {
                v11 := GenS3{F0: -113}
                var v12 int = 127
                _ = v11
                _ = v12
            } else {
                v1 = (v1 | v1)
            }
            var v13 int8 = -1
            _ = v10
            _ = v13
        }
        v14 := GenS2{F0: 0}
        _ = v14
    }
    v3 -= v3
    var v15 int = 0
    v4 = append(v4, "")
    v16 := GenS1{F0: (("foo bar" + "foo bar") + "foo bar")}
    var v17 float64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
