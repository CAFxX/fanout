package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0(p0 bool, p1 uint8) uint32 {
    v5 := []int{-293, -1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{}
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    switch ("foo bar" + "hello") {
    case "":
        v6 := GenS1{F0: 127}
        _ = v6
    case "\n\t":
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 36}
            _ = fi7
            _ = v8
        }
    case "hello":
        v1 = (v1 + v1)
    }
    switch ("hello" + "a") {
    case "\n\t":
        if (false && (-541 < 1)) {
            v1 = v1
            v3["\n\t"] = -1
            switch "hello" {
            case "a":
                var v9 int64 = v1
                _ = v9
            case "hello":
                v10 := v3[""]
                _ = v10
            default:
                var v11 int = -2147483648
                _ = v11
            }
        } else {
            _ = len(v4)
            for fi12 := 0; fi12 < 13; fi12++ {
                v1 += 1
                _ = fi12
            }
            var v13 int = 335
            _ = v13
        }
    }
    switch "" {
    case "hello":
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS2{F0: -17}
            v1 += v1
            if (((!true) && (4294967295 < 4294967295)) && (9223372036854775807 > (v1 & v1))) {
                v16 := GenS3{F0: 0, F1: 1000, F2: 33}
                for fi17 := 0; fi17 < 1; fi17++ {
                    if (true || ("\n\t" <= "\n\t")) {
                        v1 = v1
                    }
                    _ = fi17
                }
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    default:
        v18 := v4["hello"]
        _ = v18
    }
    var v19 bool = (42 == 818)
    fmt.Println(v1, v1)
    var v20 uint64 = 0
    _ = len(v3)
    v20 -= (v20 / ((v20) | 1))
    v1 -= (v1 % ((-723) | 1))
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            fmt.Println(fi22, v1, fi21)
            _ = fi22
        }
        _ = fi21
    }
    v2 = append(v2, 9223372036854775807)
    var v23 string = ""
    switch ((v23 + ((v23 + "\n\t") + (v23 + v23))) + v23) {
    case "foo bar":
        if (157 <= v20) {
            for fi24 := 0; fi24 < 10; fi24++ {
                _ = len(v2)
                _ = fi24
            }
            fmt.Println(v1)
            v25 := GenS1{F0: 44}
            _ = v25
        }
    }
    v1 = int64(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
