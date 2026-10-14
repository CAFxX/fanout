package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
func genFn0() int32 {
    v4 := GenS1{F0: 18446744073709551615}
    if false {
        v5 := []int{127, 32767}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    } else {
        var v6 int32 = 161
        if (!(false && false)) {
            if true {
                if true {
                    for fi7 := 0; fi7 < 20; fi7++ {
                        v6 += 42
                        v8 := []int{}
                        _ = fi7
                        _ = v8
                    }
                }
                v6 = (^(v6))
            }
            var v9 int16 = 32767
            _ = v9
        }
        _ = v6
    }
    _ = v4
    return 389
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{("foo bar" + "\n\t"), "\n\t", (((("a" + "\n\t") + "") + "hello") + (("foo bar" + ("\n\t" + ("a" + ("hello" + (("" + "hello") + "hello"))))) + ("\n\t" + ("" + ("\n\t" + ((("" + "a") + "\n\t") + "\n\t"))))))}
    var v3 int32 = 0
    if true {
        var v10 int64 = 0
        _ = v10
    }
    v1 = (v1 + v1)
    var v11 uint = 0
    v12 := GenS2{F0: 13}
    var v13 string = ""
    v11 += (v11 + (v11 & (v11 ^ v11)))
    fmt.Println(v13)
    var v14 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
