package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    v1 := []int{-9223372036854775808, 1, 4294967295, 32768}
    v2 := []uint32{2, 0, 594}
    var v3 int = -1
    v3 -= v3
    v2 = append(v2, 0)
    var v4 uint16 = 261
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = (v4 | (v4 * v4))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v5 uint16 = (v4 % (((v4 - 1)) | 1))
    v6 := GenS1{F0: 520}
    switch (^(v3)) {
    case -1:
        if true {
            if (371 < 4294967295) {
                fmt.Println(v4)
                fmt.Println(v5)
            }
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS1{F0: 2147483647}
                fmt.Println(v3, v5)
                _ = fi7
                _ = v8
            }
        }
    case -614:
        var v9 bool = false
        _ = v9
    case 2147483647:
        v3 -= v3
    }
    var v10 uint64 = 2
    var v11 float32 = 1.5
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 bool = (!false)
        fmt.Println(fi12)
        v4 -= v5
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v3)
    if len(v1) > 0 {
        v1[0] = (v3 + v3)
    }
    var v14 bool = (1.5 >= v11)
    for fi15 := 0; fi15 < 2; fi15++ {
        if v14 {
            if true {
                v16 := GenS2{F0: 9223372036854775807}
                v17 := GenS3{F0: -943, F1: -1}
                fmt.Println(v4, v10, fi15)
                _ = v16
                _ = v17
            }
        }
        var v18 uint8 = 1
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        if ((42 <= 127) && (-2.25 <= v11)) {
            var v20 uint16 = 32767
            _ = v20
        } else {
            v21 := GenS3{F0: 1000, F1: -1}
            var v22 uint = 1
            _ = v21
            _ = v22
        }
        if true {
            v10 = 789
        } else {
            if false {
                var v23 int = v3
                _ = v23
            } else {
                if (568 != 1000) {
                    v24 := GenS2{F0: 9223372036854775807}
                    _ = v24
                } else {
                    fmt.Println(v11, v4)
                    for fi25 := 0; fi25 < 8; fi25++ {
                        v4 += (v4 >> ((255) & 63))
                        switch (fi25 + 0) {
                        case -508:
                            switch "foo bar" {
                            case "":
                                v5 = v4
                            default:
                                var v26 uint = 373
                                _ = v26
                            }
                        }
                        _ = fi25
                    }
                    v10 = v10
                }
                v11 = v11
            }
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
