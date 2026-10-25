package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 int32 = 2147483647
    v3 := []string{"a"}
    var v4 uint64 = 764
    var v5 string = "foo bar"
    v6 := []uint32{917, uint32(v4)}
    var v7 bool = true
    var v8 float32 = (-2.25 / float32(((1e10 * 0.1) * 1e10)))
    v7 = v7
    _ = len(v1)
    v9 := GenS1{F0: 1, F1: 127, F2: (v5 + (v5 + v5))}
    var v10 uint16 = 0
    v7 = (!(v7 && true))
    if (174 >= 42) {
        v7 = v7
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v4, fi11)
            v5 = v5
            _ = fi11
        }
    } else {
        var v12 int8 = 42
        if false {
            v2 += v2
            if (uint64(v10) >= (v4 << ((802) & 63))) {
                var v13 float32 = v8
                var v14 int16 = 32767
                fmt.Println(v7)
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                v12 = v12
                if ((1.5 + v8) < v8) {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 int = (fi15 % (((-(fi15))) | 1))
                        _ = fi16
                        _ = v17
                    }
                }
                _ = fi15
            }
        } else {
            var v18 float32 = (v8 + v8)
            switch -350 {
            case 2147483647:
                v18 -= 3.14159
            case 4294967295:
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := GenS2{F0: (v12 ^ (v12 << ((38) & 63))), F1: ((float32(0.1) / v18) > 1.5)}
                    v21 := GenS1{F0: 2, F1: 118, F2: (v5 + v5)}
                    for fi22 := 0; fi22 < 16; fi22++ {
                        switch (v5 + ("hello" + ((("\n\t" + ((v5 + v5) + (v5 + (("a" + v5) + "hello")))) + ("\n\t" + ("\n\t" + "a"))) + v5))) {
                        case "foo bar":
                            var v23 bool = false
                            _ = v23
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "\n\t":
                            v18 = ((float32(1.5) * (v18 / v18)) - (-2.25 - v8))
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi22
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            v24 := GenS2{F0: v12, F1: v7}
            _ = v18
            _ = v24
        }
        _ = v12
    }
    var v25 uint32 = 0
    v26 := GenS3{F0: true, F1: 79}
    fmt.Println(v7, v2, v10)
    v6 = append(v6, uint32(v4))
    if (!v7) {
        v7 = true
        v4 += 2
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v27 := GenS4{F0: 2147483647, F1: uint(v25)}
        _ = v27
    }
    fmt.Println(v25, v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
