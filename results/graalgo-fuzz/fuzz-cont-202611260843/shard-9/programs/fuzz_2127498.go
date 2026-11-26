package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!(("" + ("hello" + "foo bar")) == ("foo bar" + "hello")))
    var v3 float64 = (0.0 * (float64((3.14159 / 1e10)) - 1.5))
    var v4 uint64 = 4294967295
    var v5 int8 = 127
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: -9223372036854775808}
    var v8 string = "\n\t"
    var v9 int8 = (v5 / ((v5) | 1))
    v2 = v2
    v10 := GenS2{F0: 0, F1: 1, F2: 0}
    v9 += 42
    var v11 int8 = v9
    if v2 {
        if true {
            v11 = v9
            if true {
                if true {
                    var v12 uint8 = 89
                    _ = v12
                }
                v13 := GenS3{F0: 65535}
                var v14 float32 = (1e10 - 0.0)
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                if ((844 <= 42) || v2) {
                    if v2 {
                        v5 += v11
                        v16 := GenS4{F0: 207, F1: 53}
                        var v17 int32 = 2147483647
                        _ = v16
                        _ = v17
                    } else {
                        v18 := []int{fi15, 9223372036854775807, (fi15 % (((fi15 + (fi15 / (((fi15 - fi15)) | 1)))) | 1)), 32767}
                        var v19 int64 = -1
                        fmt.Println(v19, fi15, v3)
                        _ = v18
                        _ = v19
                    }
                    var v20 int = fi15
                    _ = v20
                }
                switch (v8 + "") {
                case "foo bar":
                    if (!(71 >= 127)) {
                        v2 = (3.14159 <= 1e10)
                    }
                case "\n\t":
                    v2 = (!true)
                case "":
                    var v21 uint16 = 0
                    _ = v21
                }
                _ = fi15
            }
        }
        v3 += v3
        var v22 float64 = v3
        _ = v22
    } else {
        v23 := GenS2{F0: 716, F1: 4294967295, F2: 32767}
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int = 9223372036854775807
        _ = fi24
        _ = v25
    }
    v5 = (v9 ^ v5)
    var v26 float64 = 1.5
    v8 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
