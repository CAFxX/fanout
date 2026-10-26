package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint = 2
    var v2 int = 1
    var v3 int = (^(v2))
    var v4 int = v3
    var v5 int = (v3 << ((v1) & 63))
    v6 := []int{v2, 114, (v4 >> ((258) & 63)), (v5 & -9223372036854775808)}
    v7 := GenS1{F0: ("hello" + ((("" + "foo bar") + "foo bar") + "foo bar")), F1: 2, F2: v1}
    fmt.Println(v2)
    if true {
        switch (^(v4)) {
        case 705:
            v8 := GenS1{F0: "foo bar", F1: 551, F2: (^(v1))}
            _ = v8
        }
    }
    if true {
        for fi9 := 0; fi9 < 17; fi9++ {
            v5 = 32767
            _ = fi9
        }
    }
    v10 := GenS2{F0: 100, F1: 32767, F2: 42}
    var v11 float32 = 3.14159
    v3 += (v4 & (v5 & -9223372036854775808))
    fmt.Println(v3, v5, v2)
    var v12 uint64 = 2
    for fi13 := 0; fi13 < 19; fi13++ {
        if ("a" <= "hello") {
            for fi14 := 0; fi14 < 20; fi14++ {
                switch (v5 >> ((v1) & 63)) {
                case -726:
                    fmt.Println(v2, v1)
                }
                v15 := GenS1{F0: "hello", F1: 582, F2: v1}
                fmt.Println(v5, v11)
                _ = fi14
                _ = v15
            }
            if (1 <= 1) {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v6 = append(v6, (fi13 - (^(fi13))))
                    var v17 uint16 = 2
                    if len(v6) > 0 {
                        v6[0] = v5
                    }
                    _ = fi16
                    _ = v17
                }
                switch ("foo bar" + ("\n\t" + (("a" + "hello") + ("a" + "\n\t")))) {
                case "\n\t":
                    fmt.Println(v11, v3)
                case "foo bar":
                    v18 := GenS3{F0: -1, F1: ("\n\t" + "hello")}
                    _ = v18
                case "hello":
                    var v19 float32 = v11
                    _ = v19
                }
                var v20 uint32 = 512
                _ = v20
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS4{F0: -9223372036854775808}
                v23 := GenS1{F0: "foo bar", F1: 2147483648, F2: (^(v1))}
                v1 += 42
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            v5 = (-(v4))
        }
        _ = fi13
    }
    _ = len(v6)
    v24 := GenS5{F0: -32768, F1: 100}
    if true {
        v4 -= 218
        if false {
            fmt.Println(v4, v12, v5)
        }
    }
    fmt.Println(v4, v12)
    v2 = (v5 - (v4 & 42))
    var v25 uint16 = 463
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
