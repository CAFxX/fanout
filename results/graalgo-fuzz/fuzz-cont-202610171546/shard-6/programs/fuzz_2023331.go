package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    var v1 uint8 = 129
    var v2 uint64 = 2
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 -= v1
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 18; fi5++ {
            v2 = (v2 >> ((1) & 63))
            if false {
                v6 := []int{-579, fi5}
                var v7 uint = 0
                if (-1 == 1000) {
                    switch (fi4 * fi4) {
                    case 32767:
                        delete(v3, "a")
                    }
                    fmt.Println(fi5, v2)
                    v1 = uint8(v7)
                } else {
                    var v8 float32 = (0.1 / 1.0)
                    _ = len(v6)
                    v7 = v7
                    _ = v8
                }
                delete(v3, ((("" + "hello") + "\n\t") + ("\n\t" + "foo bar")))
                v1 += v1
                var v9 int16 = int16(fi4)
                delete(v3, "\n\t")
                fmt.Println(v7, v1)
                v7 += uint(v2)
                switch (-(fi4)) {
                case 2147483647:
                    v7 -= v7
                case 9223372036854775807:
                    v7 += v7
                case 32767:
                    var v10 uint16 = 32767
                    _ = v10
                }
                var v11 float32 = (3.14159 - -2.25)
                _ = v11
                _ = v6
                _ = v7
                _ = v9
            } else {
                delete(v3, "foo bar")
                fmt.Println(fi5)
            }
            var v12 int16 = 1
            _ = len(v3)
            if true {
                switch fi5 {
                case -959:
                    var v13 int16 = v12
                    _ = v13
                }
                v1 = v1
            } else {
                v14 := GenS1{F0: v1, F1: ((1 >= 4294967295) && true), F2: (fi5 | (fi5 >> ((255) & 63)))}
                switch (fi4 + fi4) {
                case 0:
                    if false {
                        var v15 uint32 = 2
                        var v16 bool = true
                        _ = len(v3)
                        _ = v15
                        _ = v16
                    } else {
                        v17 := []int{fi5, fi4, (fi5 % (((^(fi5))) | 1)), (fi5 | -2147483648)}
                        _ = v17
                    }
                case -9223372036854775808:
                    v18 := GenS1{F0: 127, F1: (!(1.5 < 0.1)), F2: fi5}
                    _ = v18
                default:
                    fmt.Println(v2, v12)
                }
                _ = v14
            }
            _ = fi5
            _ = v12
        }
        _ = fi4
    }
    if false {
        v19 := GenS1{F0: v1, F1: (!(32767 > 32767)), F2: -953}
        _ = v19
    }
    v20 := GenS2{F0: -46, F1: 0}
    var v21 float64 = ((-2.25 / 3.14159) * 1e10)
    var v22 float32 = 1e10
    v2 = (^(v2))
    if false {
        v3[""] = "\n\t"
        switch (("a" + "foo bar") + ("hello" + "hello")) {
        case "hello":
            v1 = 98
        }
        v23 := GenS2{F0: -544, F1: -791}
        _ = v23
    }
    v2 = (v2 & (v2 * v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
