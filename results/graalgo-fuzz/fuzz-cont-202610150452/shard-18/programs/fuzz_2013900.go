package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint8 {
    v5 := GenS1{F0: ("" + (("\n\t" + ("a" + "a")) + "a"))}
    v6 := []int{755}
    if false {
        switch 4294967295 {
        case 1:
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(fi7)
                _ = fi7
            }
        }
        v8 := GenS1{F0: ("foo bar" + ("\n\t" + "\n\t"))}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v10 int8 = 42
            _ = fi9
            _ = v10
        }
    }
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []string{}
    var v2 bool = ((0.0 / 3.14159) < float64(-2.25))
    v3 := make(map[string]uint32)
    v4 := []string{}
    v2 = false
    v11 := v3["hello"]
    var v12 string = ("a" + ("" + ("foo bar" + "a")))
    for fi13 := 0; fi13 < 13; fi13++ {
        _ = len(v3)
        v2 = v2
        _ = fi13
    }
    v14 := GenS1{F0: v12}
    v11 = v11
    _ = len(v4)
    fmt.Println(v2, v2)
    v15 := GenS1{F0: (v12 + (((("hello" + ((v12 + "a") + "foo bar")) + (v12 + ((v12 + v12) + v12))) + ((v12 + ((((((v12 + (v12 + "")) + v12) + v12) + "foo bar") + "a") + v12)) + v12)) + "foo bar"))}
    var v16 string = ("foo bar" + "")
    v4 = append(v4, "foo bar")
    v12 = v16
    v16 = v12
    var v17 uint8 = 255
    switch ("hello" + v12) {
    case "\n\t":
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint16 = 65535
                switch "" {
                case "a":
                    var v21 uint64 = 42
                    _ = v21
                case "\n\t":
                    var v22 bool = false
                    _ = v22
                case "":
                    var v23 int64 = -9223372036854775808
                    _ = v23
                }
                if (!true) {
                    var v24 string = (v12 + v16)
                    v25 := GenS1{F0: "hello"}
                    _ = v24
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    v12 = v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
