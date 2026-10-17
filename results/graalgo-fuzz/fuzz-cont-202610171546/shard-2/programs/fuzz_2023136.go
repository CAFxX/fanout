package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int16 }
type GenS3 struct { F0 int; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint8 = 1
    v2 := []int{423, 261, 4294967295}
    v3 := []int{14, -9223372036854775808}
    v4 := GenS1{F0: ("\n\t" + (("hello" + "a") + "")), F1: 65535}
    _ = len(v3)
    v1 = 128
    var v5 int32 = 2147483647
    if false {
        if len(v3) > 0 {
            v3[0] = -140
        }
        var v6 uint16 = 0
        v7 := GenS2{F0: 0, F1: (("" + "foo bar") + ("" + "hello")), F2: 42}
        _ = v6
        _ = v7
    } else {
        v1 = (v1 - v1)
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v1, v1, fi8)
            switch (fi8 + (fi8 % ((1) | 1))) {
            case 2147483647:
                v5 -= (-(v5))
            case 1:
                var v9 uint32 = 2147483647
                _ = v9
            case 506:
                v5 -= v5
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: ("hello" + (("foo bar" + (((("" + "") + "") + ("hello" + "foo bar")) + "\n\t")) + ((("hello" + "foo bar") + "") + ("foo bar" + ("\n\t" + "foo bar"))))), F1: 65535}
                var v12 int32 = v5
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
        }
        if (-2147483648 < 2147483647) {
            var v13 float32 = -2.25
            _ = v13
        } else {
            v14 := GenS3{F0: 127, F1: 1, F2: 32767}
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS4{F0: 2}
                v17 := GenS4{F0: 295}
                fmt.Println(v1, fi15)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
    }
    v2 = append(v2, 4294967295)
    v1 += (v1 ^ 96)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 -= (v5 & (^(v5)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
