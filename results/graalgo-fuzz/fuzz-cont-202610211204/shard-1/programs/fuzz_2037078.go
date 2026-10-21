package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 int16 = -32768
    v2 := []string{"", "hello"}
    v3 := make(map[string]bool)
    v1 -= v1
    v4 := v3["foo bar"]
    var v5 int32 = 1
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if v4 {
        if v4 {
            v4 = (true && v4)
        } else {
            v4 = true
            var v6 uint16 = 2
            v7 := GenS1{F0: v4}
            _ = v6
            _ = v7
        }
        v1 += 0
    } else {
        if len(v2) > 0 {
            v2[0] = ("a" + ((("a" + "\n\t") + "foo bar") + "a"))
        }
        v1 -= v1
    }
    v8 := GenS1{F0: v4}
    var v9 int = 32768
    if (0.1 != float64(-2.25)) {
        var v10 float64 = 3.14159
        v9 = -2147483648
        _ = v10
    } else {
        var v11 int32 = -1
        _ = v11
    }
    v1 = v1
    v9 = (v9 ^ 0)
    v3["foo bar"] = (127 > 1)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float64 = -2.25
        for fi14 := 0; fi14 < 13; fi14++ {
            switch fi12 {
            case 651:
                var v15 int = 1
                _ = v15
            case 4294967295:
                v16 := GenS1{F0: (-9223372036854775808 >= -9223372036854775808)}
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                if false {
                    fmt.Println(v9, v1)
                }
                _ = v16
            }
            if (!(-128 != -47)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch (fi14 - -1) {
                case -1:
                    v2 = append(v2, "a")
                case 0:
                    _ = len(v2)
                case 42:
                    if v4 {
                        v1 -= (-(v1))
                    } else {
                        var v17 uint64 = 1
                        _ = len(v2)
                        _ = v17
                    }
                }
                fmt.Println(v9, v13)
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v18 := GenS1{F0: (!(42 > 42))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
