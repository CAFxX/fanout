package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0() bool {
    var v7 string = "\n\t"
    v8 := GenS1{F0: 18446744073709551615}
    _ = v7
    _ = v8
    return (false && true)
}

func main() {
    v1 := []string{"hello", "foo bar", "foo bar"}
    var v2 bool = false
    var v3 int64 = 42
    var v4 bool = (!(233 < 2))
    v5 := make(map[string]int)
    var v6 bool = (!(!true))
    var v9 int16 = 0
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS1{F0: 42}
        _ = fi10
        _ = v11
    }
    v4 = (42 != -419)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            if (-1 <= -150) {
                if (4294967295 == 1) {
                    fmt.Println(v6)
                    var v14 float64 = 0.1
                    var v15 uint8 = 42
                    _ = v14
                    _ = v15
                } else {
                    v2 = false
                    fmt.Println(v4)
                    var v16 uint16 = 343
                    _ = v16
                }
                if v4 {
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v9 = v9
                    }
                    v17 := GenS2{F0: -2147483648, F1: 2147483647}
                    v6 = v6
                    _ = v17
                } else {
                    v18 := GenS3{F0: fi12, F1: 1, F2: 1}
                    v3 -= (v3 + (v3 ^ (v3 * v3)))
                    v3 = 1
                    switch (fi12 + fi13) {
                    case -2147483648:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 5:
                        _ = len(v1)
                        v6 = true
                    }
                    _ = v18
                }
            }
            fmt.Println(v2, v9)
            var v19 int16 = v9
            _ = fi13
            _ = v19
        }
        v20 := GenS3{F0: (fi12 + 127), F1: 0, F2: 1}
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (-(fi21)) {
            case 0:
                fmt.Println(fi21, v2)
            }
            var v22 uint32 = 2147483647
            _ = fi21
            _ = v22
        }
        _ = fi12
        _ = v20
    }
    v3 -= -924
    if true {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 float32 = 3.14159
            if (true || (!true)) {
                v4 = v4
            }
            _ = fi23
            _ = v24
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            v5["\n\t"] = (fi25 + fi25)
            v26 := GenS4{F0: 2, F1: 1, F2: 100}
            _ = fi25
            _ = v26
        }
        var v27 int = 42
        _ = v27
    } else {
        v9 = -795
        if v2 {
            v6 = true
        }
        if ((v9 % ((1) | 1)) <= (v9 ^ 1)) {
            fmt.Println(v9, v6, v4)
            _ = len(v5)
        } else {
            delete(v5, "a")
            switch -111 {
            case -9223372036854775808:
                for fi28 := 0; fi28 < 14; fi28++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v29 int32 = -1
                    _ = fi28
                    _ = v29
                }
            }
        }
    }
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "":
        v6 = (0 >= 588)
    }
    v30 := v5["\n\t"]
    var v31 uint64 = 362
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
