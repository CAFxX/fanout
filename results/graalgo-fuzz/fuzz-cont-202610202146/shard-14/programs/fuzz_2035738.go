package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 uint8) uint16 {
    var v6 uint = 4294967295
    var v7 uint64 = 628
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 int16 = 42
    v2 := []string{"\n\t", (("hello" + ("hello" + "a")) + ("hello" + "foo bar")), "", ("\n\t" + "")}
    var v3 uint16 = 65535
    var v4 bool = true
    v5 := make(map[string]bool)
    v2 = append(v2, ("foo bar" + ("" + ("foo bar" + ""))))
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 += (v1 * 0)
        _ = fi8
    }
    fmt.Println(v1, v3)
    v9 := GenS1{F0: 0}
    var v10 int = 4294967295
    if v4 {
        var v11 uint64 = 917
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int64 = 0
            switch (fi12 * 921) {
            case 9223372036854775807:
                var v14 int8 = 1
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        switch ("a" + "a") {
        case "\n\t":
            delete(v5, ((("a" + ("a" + "\n\t")) + "") + "hello"))
        case "foo bar":
            if ((((1e10 + 1.5) / 1.0) * float32(1e10)) > float32(-2.25)) {
                delete(v5, "foo bar")
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v15 := GenS1{F0: 4294967295}
                _ = v15
            }
        }
        _ = v11
    }
    v10 += (v10 % (((v10 / ((-572) | 1))) | 1))
    v16 := GenS2{F0: 2147483647, F1: 42, F2: 42}
    if v4 {
        v17 := GenS3{F0: 421}
        for fi18 := 0; fi18 < 5; fi18++ {
            switch v10 {
            case -348:
                var v19 bool = (18446744073709551615 == 174)
                _ = v19
            }
            _ = fi18
        }
        _ = v17
    } else {
        fmt.Println(v1)
        v20 := GenS3{F0: 255}
        fmt.Println(v3, v1)
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v22 := GenS3{F0: 255}
        v23 := GenS3{F0: 42}
        _ = fi21
        _ = v22
        _ = v23
    }
    if (!v4) {
        fmt.Println(v3)
        for fi24 := 0; fi24 < 4; fi24++ {
            v5["hello"] = false
            _ = fi24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
