package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"\n\t", "\n\t", "\n\t", "\n\t"}
    var v3 bool = true
    var v4 int8 = 0
    var v5 bool = true
    var v6 uint16 = 2
    switch "\n\t" {
    case "":
        if v5 {
            v7 := GenS1{F0: 18446744073709551615}
            _ = v7
        }
    case "hello":
        var v8 uint = 934
        _ = v8
    }
    v9 := v1[-658]
    if true {
        if true {
            fmt.Println(v6)
            if v3 {
                fmt.Println(v3)
            }
            v4 += v4
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v10 := GenS1{F0: 0}
            _ = v10
        }
    }
    v2 = append(v2, ((("foo bar" + "foo bar") + "foo bar") + ("foo bar" + "hello")))
    v3 = (v5 && (76 != 2))
    if (v9 > 2147483647) {
        var v11 int64 = 809
        _ = v11
    } else {
        v9 = 2147483647
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int8 = -128
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: 222}
    v15 := GenS1{F0: 1}
    if ((-2147483648 > -2147483648) || true) {
        v4 = (v4 << ((133) & 63))
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                v3 = v3
                v3 = (("hello" + "") <= "foo bar")
                var v18 uint16 = 65535
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                fmt.Println(v4)
                fmt.Println(v6, fi19)
                var v20 int = 32767
                _ = fi19
                _ = v20
            }
            _ = fi16
        }
        if v3 {
            if (!(((0.0 * 0.0) / (0.1 - -2.25)) < ((3.14159 / (1e10 * float64(0.1))) / -2.25))) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v4 = (v4 + v4)
            }
            if ((!v5) || (v4 <= (v4 << ((2) & 63)))) {
                v21 := GenS1{F0: 0}
                fmt.Println(v9, v4)
                _ = v21
            }
            var v22 int32 = 775
            _ = v22
        }
    }
    var v23 int8 = 127
    var v24 float64 = (-2.25 / float64((3.14159 - -2.25)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
