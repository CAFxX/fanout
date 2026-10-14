package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }

func main() {
    var v1 int8 = 21
    var v2 int32 = 349
    v3 := []string{(("hello" + "") + ("hello" + ((("\n\t" + "\n\t") + "") + ("a" + ("\n\t" + ("a" + ("\n\t" + ("\n\t" + "\n\t"))))))))}
    fmt.Println(v1, v2)
    if len(v3) > 0 {
        v3[0] = (("foo bar" + "") + "hello")
    }
    v4 := GenS1{F0: v1, F1: 0}
    fmt.Println(v2, v2)
    var v5 uint32 = 65535
    for fi6 := 0; fi6 < 18; fi6++ {
        if (2 >= 2) {
            v1 = int8(v5)
        } else {
            v7 := GenS1{F0: v1, F1: (fi6 / ((fi6) | 1))}
            var v8 float32 = 1.5
            switch "foo bar" {
            case "foo bar":
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v2, v8, v1)
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = fi9
                }
            default:
                var v10 uint32 = v5
                _ = v10
            }
            _ = v7
            _ = v8
        }
        switch "" {
        case "foo bar":
            for fi11 := 0; fi11 < 11; fi11++ {
                v1 = v1
                _ = fi11
            }
        case "":
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        switch (("" + "foo bar") + "a") {
        case "a":
            switch (-(fi6)) {
            case 42:
                if len(v3) > 0 {
                    v3[0] = (("foo bar" + "\n\t") + "")
                }
            }
        case "foo bar":
            var v12 uint = 749
            _ = v12
        case "":
            for fi13 := 0; fi13 < 20; fi13++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi13
            }
        default:
            var v14 uint32 = v5
            _ = v14
        }
        _ = fi6
    }
    v1 = 1
    if ((v2 - v2) == 2147483647) {
        if true {
            var v15 float64 = 1e10
            var v16 uint64 = 4294967295
            _ = v15
            _ = v16
        }
        v3 = append(v3, "hello")
        for fi17 := 0; fi17 < 3; fi17++ {
            if (42 > 431) {
                v18 := GenS1{F0: (v1 & v1), F1: -9223372036854775808}
                var v19 uint16 = 1
                _ = v18
                _ = v19
            }
            if (1 > 0) {
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 string = ("a" + "foo bar")
                    _ = fi20
                    _ = v21
                }
            }
            v1 = (v1 << ((210) & 63))
            _ = fi17
        }
    }
    var v22 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
