package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{}
    var v2 int16 = 32767
    var v3 uint64 = 946
    v4 := make(map[int]int)
    v5 := []string{(((("foo bar" + "foo bar") + "foo bar") + (("" + "\n\t") + ("hello" + "foo bar"))) + "foo bar"), ("hello" + (((("hello" + "") + "a") + ((((((("hello" + "") + "foo bar") + "") + "") + "a") + ("hello" + "foo bar")) + "\n\t")) + "hello")), ((("\n\t" + "a") + "a") + ("foo bar" + "a"))}
    v3 = 18446744073709551615
    v3 = (v3 * v3)
    fmt.Println(v2, v3)
    for fi6 := 0; fi6 < 18; fi6++ {
        switch fi6 {
        case 32768:
            var v7 string = ("" + "hello")
            _ = v7
        case 2147483647:
            switch (fi6 * (fi6 * -9223372036854775808)) {
            case 0:
                var v8 int16 = (v2 ^ -1)
                _ = v8
            case 32767:
                v9 := GenS1{F0: 0}
                _ = v9
            }
        case -9223372036854775808:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 bool = (711 == 0)
                v2 = 0
                var v12 uint8 = 206
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi6
    }
    var v13 int64 = -926
    if true {
        var v14 bool = (!(996 > 127))
        _ = v14
    } else {
        fmt.Println(v13)
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                if false {
                    var v17 int16 = v2
                    if true {
                        var v18 int32 = 1
                        v13 -= v13
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v19 string = ""
                    v2 = int16(fi15)
                    var v20 int = fi15
                    _ = v19
                    _ = v20
                }
                _ = fi16
            }
            v3 += v3
            _ = fi15
        }
        v21 := GenS1{F0: 18446744073709551615}
        _ = v21
    }
    if false {
        if (true || false) {
            v3 = 2
            var v22 float64 = 0.0
            if true {
                _ = len(v1)
                var v23 int64 = (v13 / ((-9223372036854775808) | 1))
                v24 := GenS1{F0: 32}
                _ = v23
                _ = v24
            }
            _ = v22
        } else {
            if false {
                var v25 bool = true
                v5 = append(v5, ("a" + "a"))
                _ = v25
            }
            v3 = (v3 - (v3 | 476))
            v26 := GenS1{F0: 1}
            _ = v26
        }
    }
    v13 = (v13 * (v13 | v13))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
