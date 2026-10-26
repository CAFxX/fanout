package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 int32; F2 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{2147483647}
    var v3 int32 = -245
    v3 -= (v3 + int32(v1))
    switch ((("foo bar" + "foo bar") + "foo bar") + "\n\t") {
    case "":
        for fi4 := 0; fi4 < 6; fi4++ {
            if (42 != 65535) {
                for fi5 := 0; fi5 < 15; fi5++ {
                    if ((128 == 193) && false) {
                        switch (("" + "") + "foo bar") {
                        case "a":
                            var v6 float64 = 0.1
                            _ = v6
                        }
                        var v7 int64 = (v1 * v1)
                        var v8 float64 = 0.0
                        _ = v7
                        _ = v8
                    }
                    _ = fi5
                }
            }
            _ = fi4
        }
    }
    _ = len(v2)
    var v9 int8 = -1
    v2 = append(v2, 42)
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = append(v2, 2)
        v1 = (v1 % ((0) | 1))
        _ = fi10
    }
    switch "foo bar" {
    case "hello":
        if (false && false) {
            if (!(!(535 > 4294967295))) {
                var v11 float64 = 1e10
                for fi12 := 0; fi12 < 14; fi12++ {
                    _ = len(v2)
                    _ = fi12
                }
                v1 = 42
                _ = v11
            }
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v1 += (v1 & v1)
            }
            var v13 int64 = (v1 * -1)
            _ = v13
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v14 := GenS1{F0: -2147483648, F1: ("hello" + "hello")}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS2{F0: 42, F1: (-(v3)), F2: 69}
        var v17 int = -2147483648
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
