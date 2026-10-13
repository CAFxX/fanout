package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int = -2147483648
    v2 := []int{128, -2147483648, v1}
    v3 := make(map[uint32]bool)
    v4 := []string{"", "foo bar", ("" + "hello"), "\n\t"}
    var v5 uint64 = 0
    v6 := []string{"\n\t"}
    v7 := GenS1{F0: 0}
    fmt.Println(v1, v5)
    v1 = v1
    v1 -= (^(v1))
    if true {
        if (true || true) {
            if (0 == 255) {
                var v8 uint = 606
                _ = v8
            }
            switch v1 {
            case 4294967295:
                v1 += (v1 * 127)
            case 0:
                v3[2] = (v5 != 2)
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if (1e10 > (1.5 / -2.25)) {
                v5 -= v5
            }
        }
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            v5 += (v5 - (v5 ^ 2))
            fmt.Println(v1)
            _ = fi9
        }
        v10 := GenS1{F0: 2}
        _ = v10
    }
    var v11 string = ("a" + "hello")
    _ = len(v3)
    v12 := GenS1{F0: 212}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        if ((!false) || ((65535 == 811) || (32767 < -233))) {
            var v14 int16 = 0
            if ((v1 >> ((255) & 63)) == v1) {
                var v15 float64 = 1.5
                if true {
                    v16 := GenS1{F0: 129}
                    switch (v11 + "\n\t") {
                    case "hello":
                        v3[500] = (966 < 2)
                    }
                    v1 += fi13
                    _ = v16
                }
                v14 = (v14 + (^(v14)))
                _ = v15
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    fmt.Println(v1, v5)
                    _ = fi17
                }
            }
            var v18 uint32 = 84
            _ = v14
            _ = v18
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                if false {
                    v5 += (v5 | v5)
                    if len(v6) > 0 {
                        v6[0] = (v11 + v11)
                    }
                } else {
                    v20 := GenS1{F0: 0}
                    var v21 uint32 = 85
                    _ = v20
                    _ = v21
                }
                _ = fi19
            }
            v1 = 32767
        }
        var v22 float64 = 0.1
        _ = fi13
        _ = v22
    }
    switch (v1 << ((4294967295) & 63)) {
    case 638:
        var v23 int8 = -1
        _ = v23
    case 127:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
