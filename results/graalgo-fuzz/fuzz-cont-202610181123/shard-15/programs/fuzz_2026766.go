package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 int; F1 int; F2 string }

func main() {
    var v1 float64 = (1.5 / 3.14159)
    var v2 uint64 = 1
    v3 := []string{"hello", ("a" + ((("foo bar" + "hello") + "hello") + "\n\t"))}
    var v4 uint16 = 42
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 float64 = v1
        var v7 int64 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 = (v1 * 0.1)
    if true {
        if false {
            var v8 uint64 = v2
            if true {
                v2 -= v8
                var v9 float32 = 0.0
                var v10 float32 = v9
                _ = v10
                _ = v9
            }
            _ = v8
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            var v13 int = fi11
            var v14 float64 = v1
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    switch "hello" {
    case "hello":
        var v15 int64 = -1
        _ = v15
    }
    v2 = (v2 | 988)
    if true {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint = 0
            v2 += 42
            fmt.Println(fi16, v2, fi16)
            _ = fi16
            _ = v17
        }
    } else {
        switch -1 {
        case 42:
            fmt.Println(v2, v4)
        }
        v4 += uint16(v2)
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 bool = (374 < 427)
        for fi20 := 0; fi20 < 7; fi20++ {
            v19 = v19
            v21 := GenS1{F0: v2, F1: (fi18 ^ 4294967295), F2: v4}
            var v22 int64 = -843
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            if v19 {
                var v24 string = ""
                _ = v24
            } else {
                var v25 bool = ((0.0 == (1.5 / (0.1 + 1.5))) || (v19 && (v19 && ((685 <= 42) || ((!true) || (-123 == 0))))))
                fmt.Println(v19, v19, fi23)
                _ = v25
            }
            v26 := GenS2{F0: 101, F1: 606}
            v2 -= (v2 + (v2 - (v2 & (v2 % (((^(v2))) | 1)))))
            _ = fi23
            _ = v26
        }
        _ = fi18
        _ = v19
    }
    switch -2147483648 {
    case 4294967295:
        if false {
            var v27 int16 = 1000
            v27 = v27
            var v28 int = 32767
            _ = v27
            _ = v28
        } else {
            v29 := GenS1{F0: uint64(v4), F1: 32768, F2: (^(v4))}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v29
        }
    case 2147483647:
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                var v32 uint16 = 65535
                fmt.Println(v32)
                _ = fi31
                _ = v32
            }
            v33 := GenS3{F0: 1, F1: 558, F2: (fi30 * 9223372036854775807)}
            switch (fi30 + 9223372036854775807) {
            case 32768:
                v34 := GenS3{F0: 18446744073709551615, F1: 1, F2: -815}
                _ = v34
            default:
                v2 += 0
            }
            _ = fi30
            _ = v33
        }
    default:
        v35 := GenS4{F0: 32768, F1: 32768, F2: ""}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
