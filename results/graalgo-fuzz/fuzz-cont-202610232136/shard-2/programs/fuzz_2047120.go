package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int = -864
    var v2 string = "\n\t"
    v3 := make(map[string]string)
    var v4 int64 = -1
    var v5 bool = (2147483647 <= 42)
    v4 = 9223372036854775807
    v4 -= 0
    v5 = ((!true) || v5)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 127
        _ = fi6
        _ = v7
    }
    var v8 int16 = -747
    var v9 int16 = (^(v8))
    v2 = ("\n\t" + ("" + "\n\t"))
    fmt.Println(v5)
    switch (v1 / ((v1) | 1)) {
    case 731:
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint = 584
            fmt.Println(v4, v11, v11)
            var v12 int32 = -1
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if false {
        if true {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS1{F0: -2147483648}
                var v15 int = -2147483648
                v16 := GenS1{F0: 596}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 int = 32767
            v2 = "\n\t"
            _ = v17
        }
    } else {
        v18 := GenS2{F0: (v4 | v4), F1: 512, F2: v4}
        _ = v18
    }
    switch (v2 + v2) {
    case "foo bar":
        for fi19 := 0; fi19 < 9; fi19++ {
            if true {
                for fi20 := 0; fi20 < 16; fi20++ {
                    if (!v5) {
                        v2 = ("" + (((("" + "hello") + v2) + (v2 + (v2 + (("\n\t" + v2) + v2)))) + v2))
                    }
                    v21 := GenS3{F0: "hello", F1: 42, F2: 32767}
                    _ = fi20
                    _ = v21
                }
                _ = len(v3)
                if (!(1 > 353)) {
                    v22 := GenS4{F0: 718}
                    var v23 uint32 = 4294967295
                    var v24 uint = 151
                    _ = v22
                    _ = v23
                    _ = v24
                }
            }
            _ = fi19
        }
    case "hello":
        v25 := GenS4{F0: 42}
        _ = v25
    }
    v26 := GenS2{F0: v4, F1: 0, F2: -9223372036854775808}
    var v27 uint8 = 1
    if v5 {
        v4 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
