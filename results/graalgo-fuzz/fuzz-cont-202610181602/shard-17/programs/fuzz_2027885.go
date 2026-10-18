package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int8 = 100
    var v2 bool = false
    var v3 uint64 = 2
    var v4 int32 = -2147483648
    v1 = v1
    switch (((("a" + (("foo bar" + ("" + "\n\t")) + "foo bar")) + "\n\t") + "\n\t") + "\n\t") {
    case "a":
        for fi5 := 0; fi5 < 7; fi5++ {
            if (80 >= (v1 & v1)) {
                switch int(v4) {
                case 206:
                    v6 := []int{fi5, -9223372036854775808, fi5}
                    _ = v6
                case 0:
                    v3 = v3
                case 1:
                    var v7 int64 = 1
                    _ = v7
                }
            }
            v8 := GenS1{F0: (v1 ^ (v1 % ((v1) | 1)))}
            var v9 int32 = v4
            _ = fi5
            _ = v8
            _ = v9
        }
    case "hello":
        if v2 {
            var v10 int16 = 32767
            v2 = (v2 && v2)
            v3 -= (v3 >> ((147) & 63))
            _ = v10
        } else {
            v11 := []int{32768}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{}
        _ = len(v13)
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    v3 = (v3 * 877)
    v14 := []int{-9223372036854775808, 4294967295}
    fmt.Println(v4, v4)
    fmt.Println(v1, v1, v1)
    switch 4294967295 {
    case -27:
        if (!((1.5 * ((1.5 * 1e10) + (1.5 * float32(1.5)))) < float32(1.5))) {
            if len(v14) > 0 {
                v14[0] = -802
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi15
            }
            switch -82 {
            case 4294967295:
                if ((3.14159 * (float64((1e10 - 3.14159)) / ((1.5 * (1e10 / 1e10)) - 3.14159))) == 0.0) {
                    v1 = 100
                    if v2 {
                        var v16 int64 = 1
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v16
                    }
                }
            }
        } else {
            var v17 uint64 = (v3 & v3)
            _ = v17
        }
    case 462:
        fmt.Println(v2, v3)
    default:
        fmt.Println(v2)
    }
    var v18 int = 128
    v19 := GenS2{F0: 42}
    var v20 uint32 = 1
    var v21 int32 = (v4 >> ((610) & 63))
    fmt.Println(v3, v4, v1)
    v22 := GenS3{F0: 1}
    v21 = (-(v21))
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
