package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []string{"a", ("\n\t" + "foo bar"), "foo bar", "a"}
    v2 := []string{}
    var v3 uint32 = 741
    var v4 string = "a"
    var v5 float64 = 0.1
    var v6 int32 = -1
    if (954 >= -9223372036854775808) {
        if len(v1) > 0 {
            v1[0] = v4
        }
    }
    fmt.Println(v5, v4)
    var v7 uint8 = 255
    fmt.Println(v3, v3)
    v7 -= v7
    switch "" {
    case "foo bar":
        var v8 string = "\n\t"
        _ = v8
    case "hello":
        switch int(v3) {
        case 32767:
            v9 := GenS1{F0: 0}
            _ = v9
        case 4294967295:
            v3 -= v3
        case 674:
            v3 = (v3 / ((52) | 1))
        }
    }
    v10 := GenS1{F0: 32768}
    if (("\n\t" != ("foo bar" + "\n\t")) || false) {
        if (!(false && false)) {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 int64 = -1
                if ((0.1 + 1.5) == v5) {
                    switch 127 {
                    case -1:
                        var v13 uint64 = uint64(v3)
                        _ = v13
                    case -527:
                        v6 -= v6
                    case 128:
                        v6 += (v6 / ((2147483647) | 1))
                    }
                }
                _ = fi11
                _ = v12
            }
            v3 = v3
        }
        if len(v2) > 0 {
            v2[0] = ((("foo bar" + v4) + "hello") + v4)
        }
    } else {
        var v14 uint16 = 840
        _ = v14
    }
    if true {
        var v15 string = v4
        _ = v15
    } else {
        if (18446744073709551615 == 675) {
            if len(v2) > 0 {
                v2[0] = v4
            }
            v7 = 146
            v16 := GenS1{F0: 127}
            _ = v16
        } else {
            if true {
                v4 = "foo bar"
            }
        }
        v6 = v6
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
