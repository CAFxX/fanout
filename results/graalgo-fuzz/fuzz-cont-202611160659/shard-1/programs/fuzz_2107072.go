package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int16; F2 uint16 }
type GenS4 struct { F0 int8; F1 int8; F2 int8 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint; F1 uint64; F2 int64 }

func main() {
    v1 := []string{("a" + ("hello" + ("" + "a"))), (("foo bar" + "foo bar") + ("a" + "foo bar")), "foo bar", "\n\t"}
    var v2 uint8 = 0
    v3 := make(map[string]int)
    var v4 uint16 = 1
    v5 := []int{}
    _ = len(v5)
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 2, F1: 0}
        v2 -= 2
        for fi8 := 0; fi8 < 4; fi8++ {
            v2 = 2
            v9 := GenS1{F0: 4294967295, F1: -1}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v10 float32 = 0.0
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v4 = (v4 << ((65535) & 63))
    fmt.Println(v4, v2)
    v4 -= v4
    if false {
        v11 := GenS2{F0: ("hello" + "\n\t")}
        v12 := GenS2{F0: "\n\t"}
        _ = v11
        _ = v12
    }
    var v13 uint = 123
    if true {
        if true {
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS3{F0: fi15, F1: 42, F2: 968}
                    v13 += 2
                    fmt.Println(fi14, v13)
                    _ = fi15
                    _ = v16
                }
                var v17 string = "a"
                _ = fi14
                _ = v17
            }
            v1 = append(v1, ("a" + ("\n\t" + "")))
            switch -2147483648 {
            case 2147483647:
                switch 9223372036854775807 {
                case 4294967295:
                    var v18 uint16 = v4
                    _ = v18
                default:
                    switch -337 {
                    case 0:
                        switch (("\n\t" + "\n\t") + "\n\t") {
                        case "hello":
                            v4 += v4
                        }
                    case 42:
                        _ = len(v3)
                    }
                }
            }
        } else {
            v19 := GenS1{F0: v13, F1: -1}
            v2 -= (v2 - v2)
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 string = "foo bar"
            switch (v21 + "hello") {
            case "foo bar":
                v4 += v4
            case "\n\t":
                var v22 bool = (false && (3.14159 == (1.5 - (1e10 * -2.25))))
                _ = v22
            case "":
                if (((42 != 2147483647) && false) && (!(true || false))) {
                    fmt.Println(v13)
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v2 -= v2
                        v24 := GenS4{F0: 55, F1: 1, F2: 70}
                        var v25 string = v21
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS5{F0: true, F1: 4294967295}
                    _ = v26
                }
            }
            v27 := GenS6{F0: 0, F1: 42, F2: -636}
            _ = fi20
            _ = v21
            _ = v27
        }
    } else {
        var v28 float32 = (0.0 + 0.0)
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
