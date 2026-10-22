package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]string)
    v3 := []string{"", "\n\t", ""}
    var v4 uint = 1
    var v5 string = "foo bar"
    if true {
        for fi6 := 0; fi6 < 20; fi6++ {
            switch (fi6 | (fi6 % ((fi6) | 1))) {
            case 32767:
                v1 = (v1 >> ((v4) & 63))
            default:
                var v7 int64 = -1
                _ = v7
            }
            _ = fi6
        }
        fmt.Println(v1, v5, v1)
        v5 = (v5 + "\n\t")
    }
    switch ("a" + v5) {
    case "a":
        for fi8 := 0; fi8 < 9; fi8++ {
            delete(v2, (fi8 + 1))
            v3 = append(v3, "")
            _ = fi8
        }
    }
    if (v4 < (v4 % ((v4) | 1))) {
        if (248 != 1) {
            switch "\n\t" {
            case "hello":
                var v9 int64 = -1
                _ = v9
            case "foo bar":
                fmt.Println(v5)
            case "a":
                var v10 uint64 = 4294967295
                _ = v10
            }
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(v5)
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int64 = 42
                    fmt.Println(v5, v13, v1)
                    if false {
                        var v14 float32 = 1e10
                        v4 = (v4 ^ (v4 - 335))
                        var v15 uint64 = 467
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                fmt.Println(v5, v4)
                _ = fi11
            }
        }
    } else {
        fmt.Println(v1, v4, v4)
    }
    fmt.Println(v1, v5, v4)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            if ((!false) || (!((!(!(!false))) || false))) {
                var v18 float32 = 1e10
                v18 -= float32((float64(v18) + 3.14159))
                _ = v18
            }
            v5 = ""
            _ = fi17
        }
        _ = fi16
    }
    if (!(4294967295 != -2147483648)) {
        v19 := v2[42]
        if false {
            v4 += 0
            if true {
                v19 = ((v5 + (v19 + "a")) + (v5 + v19))
                _ = len(v3)
            } else {
                var v20 int = 128
                _ = v20
            }
        } else {
            v21 := GenS1{F0: 739}
            v5 = ("foo bar" + v19)
            v22 := v2[38]
            _ = v21
            _ = v22
        }
        if true {
            var v23 uint64 = 83
            v3 = append(v3, "foo bar")
            v24 := GenS2{F0: -880}
            _ = v23
            _ = v24
        } else {
            if ((-32768 == 0) || false) {
                v25 := GenS2{F0: 4294967295}
                _ = v25
            }
        }
        _ = v19
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v26 := v2[626]
        _ = v26
    }
    _ = len(v3)
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v5)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
