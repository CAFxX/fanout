package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{}
    v2 := []string{"", "", "foo bar"}
    var v3 int8 = 0
    for fi4 := 0; fi4 < 5; fi4++ {
        switch ((("foo bar" + "foo bar") + "hello") + "hello") {
        case "foo bar":
            for fi5 := 0; fi5 < 1; fi5++ {
                _ = len(v2)
                v3 = (v3 << ((861) & 63))
                _ = fi5
            }
        case "\n\t":
            var v6 float32 = -2.25
            _ = v6
        case "hello":
            v3 = 42
        default:
            for fi7 := 0; fi7 < 2; fi7++ {
                if (-1 == 9223372036854775807) {
                    var v8 uint64 = 18446744073709551615
                    _ = v8
                }
                v3 = (v3 | v3)
                var v9 uint64 = 691
                _ = fi7
                _ = v9
            }
        }
        fmt.Println(v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi4
    }
    v10 := GenS1{F0: true}
    v3 -= v3
    var v11 uint16 = 543
    v11 -= (v11 & (v11 - v11))
    switch "a" {
    case "foo bar":
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int32 = 42
                v14 = (v14 << ((255) & 63))
                _ = fi13
                _ = v14
            }
            if false {
                v15 := GenS1{F0: ((("hello" + "\n\t") >= "a") && false)}
                _ = v15
            }
            _ = fi12
        }
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            switch "foo bar" {
            case "\n\t":
                if (!(2147483647 == (fi16 + (-(fi16))))) {
                    var v17 uint64 = 279
                    _ = v17
                }
            case "hello":
                var v18 int32 = 1
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                if ((v3 >> ((18446744073709551615) & 63)) >= v3) {
                    v3 -= (^(v3))
                    v3 += v3
                }
                _ = len(v2)
                _ = fi19
            }
            _ = fi16
        }
    default:
        v20 := GenS1{F0: false}
        _ = v20
    }
    var v21 int8 = v3
    switch 42 {
    case 127:
        v22 := GenS1{F0: (2147483648 >= 2)}
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v11 = 0
        v21 = (v3 * 0)
        var v24 int = fi23
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
}
