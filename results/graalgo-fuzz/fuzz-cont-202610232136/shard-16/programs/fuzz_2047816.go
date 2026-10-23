package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint32 = 378
    v2 := []string{"foo bar", "", "foo bar", "a"}
    var v3 string = ""
    var v4 bool = (677 >= 2147483647)
    fmt.Println(v1, v1, v4)
    v3 = ((("" + v3) + v3) + "")
    var v5 uint64 = 4294967295
    v6 := GenS1{F0: 32767, F1: true, F2: v3}
    fmt.Println(v3, v4)
    v7 := GenS1{F0: -744, F1: (529 <= 475), F2: ((v3 + "a") + "hello")}
    v8 := GenS2{F0: (v1 >= 2147483647)}
    v9 := GenS3{F0: v1, F1: (^(v5))}
    if v4 {
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                v4 = v4
                var v12 int16 = -1
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint8 = 1
                switch "hello" {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                case "\n\t":
                    v15 += (v15 - 127)
                    if false {
                        var v16 uint32 = v1
                        _ = v16
                    }
                    switch (v3 + "a") {
                    case "foo bar":
                        v17 := GenS1{F0: 395, F1: true, F2: ((("foo bar" + v3) + v3) + "hello")}
                        _ = v17
                    }
                }
                v15 += uint8(v5)
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
