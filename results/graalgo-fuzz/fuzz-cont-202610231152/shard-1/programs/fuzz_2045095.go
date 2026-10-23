package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 string = ""
    v3 := make(map[uint32]int)
    var v4 int8 = 127
    var v5 uint32 = (v1 % (((v1 * 39)) | 1))
    var v6 float64 = 0.1
    switch 716 {
    case 2147483647:
        fmt.Println(v2, v2)
    case -774:
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 float64 = float64((3.14159 * (3.14159 + (1e10 / float32(v6)))))
            var v9 int8 = (v4 >> ((31) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
    case 128:
        if true {
            if false {
                var v10 bool = false
                _ = v10
            }
            v1 = (^(v1))
            if false {
                var v11 int64 = 9223372036854775807
                var v12 uint64 = 92
                for fi13 := 0; fi13 < 15; fi13++ {
                    if true {
                        v14 := GenS1{F0: 324}
                        _ = v14
                    }
                    _ = fi13
                }
                _ = v11
                _ = v12
            }
        }
    default:
        for fi15 := 0; fi15 < 16; fi15++ {
            if false {
                var v16 float32 = 1e10
                v17 := GenS1{F0: 65535}
                _ = v16
                _ = v17
            } else {
                v4 -= v4
            }
            _ = fi15
        }
    }
    v6 += 1.5
    var v18 bool = (true || true)
    switch -193 {
    case -9223372036854775808:
        v19 := []int{-76, 464}
        _ = v19
    }
    v6 = v6
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: v5}
        for fi22 := 0; fi22 < 11; fi22++ {
            if true {
                var v23 int8 = 100
                v23 = v23
                _ = v23
            }
            switch ("foo bar" + v2) {
            case "hello":
                v1 -= v1
            case "\n\t":
                switch (((v2 + ("\n\t" + v2)) + "a") + v2) {
                case "a":
                    _ = len(v3)
                case "\n\t":
                    v24 := GenS1{F0: v5}
                    _ = v24
                }
            case "a":
                fmt.Println(fi20, v4, v2)
                var v25 float64 = -2.25
                _ = v25
            }
            _ = fi22
        }
        switch (("foo bar" + ("\n\t" + ((v2 + "\n\t") + ""))) + (v2 + "foo bar")) {
        case "foo bar":
            v18 = (v1 != v5)
        }
        _ = fi20
        _ = v21
    }
    v26 := []int{9223372036854775807, -1, 42, -9223372036854775808}
    var v27 uint32 = 2
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
