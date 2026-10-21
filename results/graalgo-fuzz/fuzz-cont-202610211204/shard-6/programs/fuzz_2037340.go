package main

import (
    "fmt"
)

func genFn0(p0 int8) bool {
    v7 := []int{}
    p0 = (p0 << ((255) & 63))
    _ = p0
    _ = v7
    return (!(0.1 <= -2.25))
}
func genFn1(p0 int16, p1 uint8, p2 int) int {
    if (65535 >= 65535) {
        for fi8 := 0; fi8 < 12; fi8++ {
            p0 = p0
            _ = fi8
        }
        p2 = (p2 & 2147483647)
        var v9 bool = (p1 <= (^(p1)))
        _ = v9
    }
    v10 := []int{(p2 & p2), -926, 32768, 997}
    var v11 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p2 % ((p2) | 1))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (1 > -1)
    var v3 string = "a"
    var v4 string = "\n\t"
    var v5 int = -1
    v6 := make(map[string]bool)
    v4 = ("a" + (("foo bar" + v4) + v4))
    if (0 > 1) {
        var v12 int16 = 1000
        _ = v12
    } else {
        v4 = ((v3 + v3) + (v4 + (v4 + "")))
        v5 = v5
        v5 = -9223372036854775808
    }
    v3 = (("foo bar" + v4) + ((("hello" + (v3 + v3)) + "hello") + (v3 + (v3 + (v4 + "")))))
    var v13 int8 = 79
    var v14 bool = (!(100 == (^(v13))))
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            if v14 {
                if ((!(1000 > 32767)) && v14) {
                    v4 = (v4 + ("\n\t" + v3))
                    v6[(v4 + "\n\t")] = v2
                } else {
                    fmt.Println(v3)
                    v2 = v2
                    v13 = v13
                }
            }
            _ = fi16
        }
        fmt.Println(v3, v13, v14)
        _ = fi15
    }
    var v17 uint64 = 2
    for fi18 := 0; fi18 < 7; fi18++ {
        v13 = (v13 & v13)
        _ = fi18
    }
    v19 := v1[-321]
    fmt.Println(v19, v14)
    if v2 {
        v3 = v3
        v20 := v6[("a" + v3)]
        _ = v20
    }
    switch genFn1(1000, 127, v5) {
    case 1:
        switch v4 {
        case "\n\t":
            v14 = false
        case "foo bar":
            fmt.Println(v19, v17)
        }
    case 703:
        v21 := []int{(v5 ^ -760), 869}
        _ = v21
    case -2147483648:
        for fi22 := 0; fi22 < 11; fi22++ {
            v5 = genFn1(1000, 255, int(v13))
            var v23 uint64 = 694
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        switch "hello" {
        case "foo bar":
            var v25 string = "hello"
            _ = v25
        case "\n\t":
            v19 = (v19 << ((215) & 63))
        default:
            v2 = false
        }
        _ = fi24
    }
    var v26 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
