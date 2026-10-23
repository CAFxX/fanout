package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    var v1 int = 42
    var v2 int64 = -399
    var v3 int64 = v2
    v2 -= v2
    if (4294967295 >= 4294967295) {
        var v4 int = v1
        v5 := []int{(^(v1)), (v4 >> ((2) & 63)), 0, (v4 - v1)}
        var v6 int = (v1 >> ((1) & 63))
        _ = v4
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := []int{v1, fi7}
            v2 = v3
            _ = fi7
            _ = v8
        }
        var v9 uint = 4294967295
        _ = v9
    }
    switch (^(v1)) {
    case 9223372036854775807:
        v1 = -1
    }
    v10 := GenS1{F0: 42}
    v11 := []int{v1, v1}
    switch (v1 * 9223372036854775807) {
    case 716:
        v12 := GenS1{F0: 2}
        _ = v12
    case -2147483648:
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int16 = -32768
            _ = fi13
            _ = v14
        }
    }
    v2 = 1
    v15 := GenS2{F0: 42, F1: v2}
    fmt.Println(v2, v2, v2)
    switch (v1 % (((v1 % (((v1 << ((1) & 63))) | 1))) | 1)) {
    case -827:
        for fi16 := 0; fi16 < 14; fi16++ {
            if true {
                v3 -= (v2 ^ (v3 / ((v2) | 1)))
            } else {
                var v17 uint8 = 164
                if len(v11) > 0 {
                    v11[0] = fi16
                }
                if false {
                    fmt.Println(v17, v3)
                    _ = len(v11)
                } else {
                    var v18 int16 = 1000
                    if false {
                        var v19 uint = 255
                        v3 = v3
                        _ = v19
                    }
                    _ = v18
                }
                _ = v17
            }
            _ = fi16
        }
    case 32767:
        switch "a" {
        case "":
            switch 42 {
            case 0:
                v3 = (v2 << ((360) & 63))
            case 128:
                v11 = append(v11, v1)
            case -2147483648:
                switch "a" {
                case "":
                    v20 := GenS1{F0: 931}
                    _ = v20
                case "hello":
                    var v21 int64 = 9223372036854775807
                    _ = v21
                }
            }
        case "hello":
            v3 -= v3
        case "foo bar":
            v2 -= (-(v2))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
